package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"

	. "github.com/stefanhans/golang-contexting/ctx"
)

type Contexting struct {
	DocId string `firestore:"doc-id"`
	Cxt   string `firestore:"cxt"`
}

func main() {

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx := context.Background()

	sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err := firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()

	collection := "context"

	//cxt := []byte{1,1,1}
	////data := []byte{0}
	//
	//fmt.Printf("%q\n", cxt)
	//fmt.Printf("%v\n", len(cxt))
	//
	//doc := client.Collection(collection).NewDoc()
	//
	//_, err = doc.Create(ctx, Contexting{
	//	DocId:       doc.ID,
	//	Cxt:        base64.StdEncoding.EncodeToString(cxt),
	//})
	//if err != nil {
	//	fmt.Printf("failed to create document: %v\n", err)
	//}

	bytesFilter := []byte{1, 1, 0}

	var bytesCxt []byte
	_ = bytesCxt

	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Printf("%v\n", doc.Data())
		var data Contexting
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}

		bytesCxt, err = base64.StdEncoding.DecodeString(data.Cxt)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if bytesFilter[0] == bytesCxt[0] && bytesFilter[1] == bytesCxt[1] && bytesFilter[2] == bytesCxt[2] {

			fmt.Printf("bytesCxt: %q\n", bytesCxt)
		}
	}

	//str := base32.StdEncoding.EncodeToString(data)
	//fmt.Printf("byte base32: %q\n", str)

	//str := base64.StdEncoding.EncodeToString(data)
	//fmt.Printf("byte base64: %q\n", str)
	//
	//str = "AAEAAQEBAAEBAQEBAQABAQABAQEAAQEAAQEBAAEBAAEBAQABAQAAAAAAAQEAAQEBAAEBAAEBAQABAQABAQEAAQEAAQEBAAEBAAEBAQABAQABAQEAAQAAAA=="
	//data, err = base64.StdEncoding.DecodeString(str)
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//fmt.Printf("%q\n", data)

	fmt.Printf("CI_BRICK_RZV: %q\n", CI_BRICK_RZV)
	fmt.Printf("CI_BRICK_RZV.Content: %q\n", CI_BRICK_RZV.Content)
	fmt.Printf("CI_BRICK_RZV.Mask: %q\n", CI_BRICK_RZV.Mask)

	ciRequest := CiBrick{
		Content: 2,
		Mask:    0,
	}

	fmt.Printf("ciRequest: %q\n", ciRequest)
	fmt.Printf("ciRequest.Content: %q\n", ciRequest.Content)
	fmt.Printf("ciRequest.Mask: %q\n", ciRequest.Mask)
}
