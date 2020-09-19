package main

import (
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func listdocs(arguments []string) {

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	iter := client.Collection("bookings").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Ref.Path)
	}
}
