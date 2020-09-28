package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/commandline/types"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
)

func AddUser(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "users"

	doc := client.Collection(collection).NewDoc()

	wr, err := doc.Create(ctx, types.User{
		DocId:       doc.ID,
		Name:        "Tesla Deluxe",
		Type:        "eCar",
		Status:      "ready",
		Description: "just a test description",
	})
	if err != nil {
		fmt.Printf("failed to create document: %v\n", err)
	}
	fmt.Println(wr.UpdateTime)
}

func StoreUsers(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "users"

	jsonUsers, err := ioutil.ReadFile("data/users.json")
	if err != nil {
		fmt.Printf("could not read file %q: %v\n", "data/users.json", err)
		return
	}

	fmt.Printf(string(jsonUsers))

	// Declared an empty map interface
	var users []types.User

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal([]byte(jsonUsers), &users)
	if err != nil {
		fmt.Printf("failed to unmarshal 'users': %v\n", err)
	}

	fmt.Printf("users: %v\n", users)

	for i, user := range users {

		doc := client.Collection(collection).NewDoc()
		user.DocId = doc.ID

		wr, err := doc.Create(ctx, user)
		if err != nil {
			fmt.Printf("failed to create document #%v: %v\n", i, err)
		}
		fmt.Println(wr.UpdateTime)
	}
}

func ClearUsers(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "users"

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

func GetUsers(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "users"

	users := make([]types.User, 0)

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

		var data types.User
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}
		users = append(users, data)
	}
	fmt.Printf("users: \n%v\n", users)

	jsonUsers, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshall 'userRequest': %v\n", err)
		return
	}
	fmt.Printf("userRequest: %s\n", jsonUsers)

}
