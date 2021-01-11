package main

import (
	"google.golang.org/api/iterator"
	"log"
	"time"

	//"encoding/json"
	"fmt"
)

func QueryBookings(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	collection = "bookings"
	from := time.Date(2020, time.September, 20, 0, 0, 0, 0, time.UTC)
	to := time.Date(2020, time.December, 21, 23, 0, 0, 0, time.UTC)

	query := client.Collection(collection).
		Where("from", ">=", from)

	iter := query.Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		//fmt.Printf("where from: %q\n\n", doc.Data())
		if doc.Data()["to"].(time.Time).Unix() <= to.Unix() {

			fmt.Printf("between yesterday and tomorrow: %q\n\n", doc.Data())
		} else {

			fmt.Printf("NOT: %q\n\n", doc.Data())
		}
	}
}
