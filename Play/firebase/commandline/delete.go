package main

import (
	"fmt"
	"log"

	"google.golang.org/api/iterator"
)

func delete(arguments []string) {

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	for _, id := range arguments {
		_, err := client.Collection(collection).Doc(id).Delete(ctx)
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}
		fmt.Printf("Document with id %q deleted\n", id)

	}
}

func deleteall(arguments []string) {

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

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
			fmt.Printf("all documents deleted\n")
			return
		}

		result, err := batch.Commit(ctx)
		if err != nil {
			fmt.Printf("failed to commit batch for deletion: %v\n", err)
			return
		}
		fmt.Printf("Commit Result: %v\n", result)

	}
}
