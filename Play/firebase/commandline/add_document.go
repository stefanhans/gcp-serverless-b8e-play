package main

import (
	//"cloud.google.com/go/firestore"
	//"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func addDocument(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	if collection == "" {
		fmt.Printf("no collection set - use %q\n", "setcollection")
		return
	}

	//document := `{
	//  "name": `+string(firestore.ServerTimestamp)+`,
	//  "country": 123,
	//  "city": "Wuhan",
	//  "reason": "Non vedge Food"
	//}`

	type State struct {
		Capital string    `firestore:"capital"`
		Now     time.Time `firestore:"pop"` // in millions
	}
	//client.Collection(collection).NewDoc().

	wr, err := client.Collection(collection).NewDoc().Create(ctx, State{
		Capital: "Denver",
		Now:     time.Now(),
	})
	if err != nil {
		fmt.Printf("failed to create document: %v\n", err)
	}
	fmt.Println(wr.UpdateTime)

	//// Declared an empty map interface
	//var result map[string]interface{}
	//
	//// Unmarshal or Decode the JSON to the interface.
	//json.Unmarshal([]byte(string(document)), &result)

	//_, _, err = client.Collection(collection).Add(ctx, wr)
	//if err != nil {
	//	log.Fatalf("Failed adding document to collection %q: %v", collection, err)
	//}

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
