package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/functions/types"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
)

var (
	client *firestore.Client
	ctx    context.Context
	//err    error
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
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	// Todo: Check for conflicts, e.g. same name

	// Check for predefined vehicle resp. document id
	if vehicleRequest.DocId == "" {

		doc := client.Collection("vehicles").NewDoc()

		// Set the document id
		vehicleRequest.DocId = doc.ID

		// Set the registration status
		vehicleRequest.Status = "registered"

		wr, err := doc.Create(ctx, vehicleRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to create document: %s", err), http.StatusInternalServerError)
			return
		}
		_ = wr
		//fmt.Println(wr.UpdateTime)

	} else {

		doc := client.Doc("vehicles/" + vehicleRequest.DocId)
		writeResult, err := doc.Create(ctx, vehicleRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to create vehicles/%s: %s", vehicleRequest.DocId, err), http.StatusInternalServerError)
			return
		}
		_ = writeResult

	}

	// Marshal vehicle
	jsonVehicle, err := json.Marshal(vehicleRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal 'vehicleRequest': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("vehicleReply: %s\n", jsonVehicle)

	// Response
	_, err = fmt.Fprint(w, string(jsonVehicle))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func GetVehicleById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("GetVehicleById Request: %s\n", body)

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
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	// "49KvH6m2Y7flEegBdPv4"
	q := client.Collection("vehicles").
		Where("DocId", "==", vehicleRequest.DocId).
		Limit(1)
	iter := q.Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		http.Error(w, fmt.Sprintf("no document found: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(doc.Data())
	vehicle := doc.Data()

	fmt.Printf("vehicle: \n%v\n", vehicle)

	jsonVehicle, err := json.MarshalIndent(vehicle, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonVehicle': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("jsonVehicle: %s\n", jsonVehicle)

	// Response
	_, err = fmt.Fprint(w, string(jsonVehicle))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func DeleteVehicleById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("DeleteVehicleById Request: %s\n", body)

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
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	deleteResult, err := client.Batch().
		Delete(client.Doc("vehicles/"+vehicleRequest.DocId), firestore.Exists).
		Commit(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete vehicle %q: %s", vehicleRequest.DocId, err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("C: \n%v\n", deleteResult)
	vehicleRequest.Description = "deleteResult: \n"
	for i, s := range deleteResult {
		vehicleRequest.Description += fmt.Sprintf("%v: %q\n", i, s)
	}

	// Marshal vehicle
	jsonVehicle, err := json.Marshal(vehicleRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal 'vehicleRequest': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("vehicleReply: %s\n", jsonVehicle)

	// Response
	_, err = fmt.Fprint(w, string(jsonVehicle))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func ClearVehicles(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClearVehicles Request: %s\n", body)

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
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
				http.Error(w, fmt.Sprintf("failed next iteration: %s", err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprintf("failed to commit batch for deletion: %s", err), http.StatusInternalServerError)
			return
		}
		_ = result
		//fmt.Printf("Commit Result: %v\n", result)

		// Response
		_, err = fmt.Fprintf(w, "{%q: %q}", "result", "cleared")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
		}
	}
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("GetVehicles Request: %s\n", body)

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprintf("failed next iteration: %s", err), http.StatusInternalServerError)
			return
		}
		fmt.Println(doc.Data())

		var data types.Vehicle
		if err := doc.DataTo(&data); err != nil {
			http.Error(w, fmt.Sprintf("failed to convert data: %s", err), http.StatusInternalServerError)
			return
		}
		vehicles = append(vehicles, data)
	}
	fmt.Printf("vehicles: \n%v\n", vehicles)

	jsonVehicles, err := json.MarshalIndent(vehicles, "    ", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonVehicles': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("vehicleRequest: %s\n", jsonVehicles)

	// Response
	_, err = fmt.Fprintf(w, "{ \n    %q: %s\n}\n", "Vehicles", string(jsonVehicles))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

//func StoreVehicles(arguments []string) {
//
//	_ = arguments
//
//	if client == nil {
//		fmt.Printf("not connected - use %q\n", "connect")
//		return
//	}
//
//	jsonVehicles, err := ioutil.ReadFile("data/vehicles.json")
//	if err != nil {
//		fmt.Printf("could not read file %q: %v\n", "data/vehicles.json", err)
//		return
//	}
//
//	fmt.Printf(string(jsonVehicles))
//
//	// Declared an empty map interface
//	var vehicles []types.Vehicle
//
//	// Unmarshal or Decode the JSON to the interface.
//	err = json.Unmarshal([]byte(jsonVehicles), &vehicles)
//	if err != nil {
//		fmt.Printf("failed to unmarshal 'vehicles': %v\n", err)
//	}
//
//	fmt.Printf("vehicles: %v\n", vehicles)
//
//	for i, vehicle := range vehicles {
//
//		doc := client.Collection("vehicles").NewDoc()
//		vehicle.DocId = doc.ID
//
//		wr, err := doc.Create(ctx, vehicle)
//		if err != nil {
//			fmt.Printf("failed to create document #%v: %v\n", i, err)
//		}
//		fmt.Println(wr.UpdateTime)
//	}
//}
