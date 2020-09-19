package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/commandline/types"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/type/latlng"
	"io/ioutil"
	"log"
)

func AddVehicle(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "vehicles"

	doc := client.Collection(collection).NewDoc()

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
}

func StoreVehicles(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "vehicles"

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

		doc := client.Collection(collection).NewDoc()
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

	collection = "vehicles"

	ref := client.Collection(collection)
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
			fmt.Printf("all %s deleted\n", collection)
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

	collection = "vehicles"

	vehicles := make([]types.Vehicle, 0)

	iter := client.Collection(collection).Documents(ctx)
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
