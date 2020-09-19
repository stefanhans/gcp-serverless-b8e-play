package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/commandline/types"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/type/latlng"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

var (
	client *firestore.Client
	ctx    context.Context
	err    error
)

func AddVehicle(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("AddVehicle Request: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var vehicleRequest types.Vehicle
	err = json.Unmarshal(bytes, &vehicleRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshall JSON input: %s", err), http.StatusInternalServerError)
		return
	}

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	doc := client.Collection("vehicles").NewDoc()

	wr, err := doc.Create(ctx, types.Vehicle{
		DocId:       doc.ID,
		Name:        "Tesla Deluxe",
		Type:        "eCar",
		Status:      "ready",
		ParkingLot:  "my garage",
		GeoPoint:    &latlng.LatLng{Latitude: 0.1, Longitude: 0.1},
		Description: "just a test description",
	})
	if err != nil {
		fmt.Printf("failed to create document: %v\n", err)
	}
	fmt.Println(wr.UpdateTime)

	// Fake the booking procedure as successful
	vehicleRequest.Status = "confirmed"

	// Marshal vehicle
	vehicleJson, err := json.Marshal(vehicleRequest)
	if err != nil {
		fmt.Printf("failed to marshall 'vehicleRequest': %v\n", err)
		return
	}
	fmt.Printf("vehicleReply: %s\n", vehicleJson)

	// Response registration
	_, err = fmt.Fprint(w, string(vehicleJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func StoreVehicles(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	jsonVehicles, err := ioutil.ReadFile("data/vehicles.json")
	if err != nil {
		fmt.Printf("could not read file %q: %v\n", "data/vehicles.json", err)
		return
	}

	fmt.Printf(string(jsonVehicles))

	// Declared an empty map interface
	var vehicles []types.Vehicle

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal([]byte(jsonVehicles), &vehicles)
	if err != nil {
		fmt.Printf("failed to unmarshal 'vehicles': %v\n", err)
	}

	fmt.Printf("vehicles: %v\n", vehicles)

	for i, vehicle := range vehicles {

		doc := client.Collection("vehicles").NewDoc()
		vehicle.DocId = doc.ID

		wr, err := doc.Create(ctx, vehicle)
		if err != nil {
			fmt.Printf("failed to create document #%v: %v\n", i, err)
		}
		fmt.Println(wr.UpdateTime)
	}
}

func ClearVehicles(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	ref := client.Collection("vehicles")
	batchSize := 20

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("failed next iteration: %v\n", err)
				return
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			fmt.Printf("all %s deleted\n", "vehicles")
			return
		}

		result, err := batch.Commit(ctx)
		if err != nil {
			fmt.Printf("failed to commit batch for deletion: %v\n", err)
			return
		}
		_ = result
		//fmt.Printf("Commit Result: %v\n", result)
	}
}

func GetVehicles(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	vehicles := make([]types.Vehicle, 0)

	iter := client.Collection("vehicles").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())

		var data types.Vehicle
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}
		vehicles = append(vehicles, data)
	}
	fmt.Printf("vehicles: \n%v\n", vehicles)

	jsonVehicles, err := json.MarshalIndent(vehicles, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshall 'vehicleRequest': %v\n", err)
		return
	}
	fmt.Printf("vehicleRequest: %s\n", jsonVehicles)

}
