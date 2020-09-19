package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/booking/types"
	"google.golang.org/api/iterator"
	"log"
)

func AddItem(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "items"

	item := types.Item{ID: 1,
		Name: "Tesla Deluxe",
		Type: types.CAR,
	}

	jsonRequest, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshall 'jsonRequest': %v\n", err)
		return
	}
	fmt.Printf("jsonRequest: %s\n", jsonRequest)

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(string(jsonRequest)), &result)

	// [START fs_add_data_1]
	_, _, err = client.Collection(collection).Add(ctx, result)
	if err != nil {
		log.Fatalf("Failed adding document to collection %q: %v", collection, err)
	}

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
	}
}

func DeleteItems(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "items"

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

func GetItems(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "items"

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
	}
}
