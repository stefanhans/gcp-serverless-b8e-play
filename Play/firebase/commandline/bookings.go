package main

import (
	"encoding/json"
	"io/ioutil"

	//"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/commandline/types"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func AddBooking(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "bookings"

	doc := client.Collection(collection).NewDoc()

	wr, err := doc.Create(ctx, types.Booking{
		DocId:         doc.ID,
		User:          "Alice",
		Vehicle:       "Tesla Deluxe",
		VehicleType:   "eCar",
		VehicleStatus: "",
		ParkingLot:    "",
		From:          time.Now(),
		To:            time.Now().Add(time.Hour * 2),
		Status:        "request",
		StatusTime:    time.Now(),
	})
	if err != nil {
		fmt.Printf("failed to create document: %v\n", err)
	}
	fmt.Println(wr.UpdateTime)

	//jsonRequest, err := json.MarshalIndent(item, "", "  ")
	//if err != nil {
	//	fmt.Printf("failed to marshall 'jsonRequest': %v\n", err)
	//	return
	//}
	//fmt.Printf("jsonRequest: %s\n", jsonRequest)
	//
	//// Declared an empty map interface
	//var result map[string]interface{}
	//
	//// Unmarshal or Decode the JSON to the interface.
	//json.Unmarshal([]byte(string(jsonRequest)), &result)
	//
	//// [START fs_add_data_1]
	//_, _, err = client.Collection(collection).Add(ctx, result)
	//if err != nil {
	//	log.Fatalf("Failed adding document to collection %q: %v", collection, err)
	//}
	//
	//iter := client.Collection(collection).Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}
}

func StoreBookings(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "bookings"

	jsonBookings, err := ioutil.ReadFile("data/bookings.json")
	if err != nil {
		fmt.Printf("could not read file %q: %v\n", "data/bookings.json", err)
		return
	}

	fmt.Printf(string(jsonBookings))

	// Declared an empty map interface
	var bookings []types.Booking

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal([]byte(jsonBookings), &bookings)
	if err != nil {
		fmt.Printf("failed to unmarshal 'bookings': %v\n", err)
	}

	fmt.Printf("bookings: %v\n", bookings)

	for i, booking := range bookings {

		doc := client.Collection(collection).NewDoc()
		booking.DocId = doc.ID

		wr, err := doc.Create(ctx, booking)
		if err != nil {
			fmt.Printf("failed to create document #%v: %v\n", i, err)
		}
		fmt.Println(wr.UpdateTime)

	}

	//jsonRequest, err := json.MarshalIndent(item, "", "  ")
	//if err != nil {
	//	fmt.Printf("failed to marshall 'jsonRequest': %v\n", err)
	//	return
	//}
	//fmt.Printf("jsonRequest: %s\n", jsonRequest)
	//
	//// Declared an empty map interface
	//var result map[string]interface{}
	//
	//// Unmarshal or Decode the JSON to the interface.
	//json.Unmarshal([]byte(string(jsonRequest)), &result)
	//
	//// [START fs_add_data_1]
	//_, _, err = client.Collection(collection).Add(ctx, result)
	//if err != nil {
	//	log.Fatalf("Failed adding document to collection %q: %v", collection, err)
	//}
	//
	//iter := client.Collection(collection).Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}
}

func ClearBookings(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "bookings"

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

func GetBookings(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "bookings"

	bookings := make([]types.Booking, 0)

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

		var data types.Booking
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}
		bookings = append(bookings, data)
	}
	fmt.Printf("bookings: \n%v\n", bookings)

	jsonBookings, err := json.MarshalIndent(bookings, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshall 'bookingRequest': %v\n", err)
		return
	}
	fmt.Printf("bookingRequest: %s\n", jsonBookings)

}
