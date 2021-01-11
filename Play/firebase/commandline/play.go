package main

import (
	"cloud.google.com/go/firestore"
	"encoding/base64"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"time"

	. "github.com/stefanhans/golang-contexting/ctx"
)

func play(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	ciRequest := CiBrick{
		Content: 6,
		Mask:    0,
	}

	fmt.Printf("ciRequest: %q\n", ciRequest)
	fmt.Printf("ciRequest.Content: %q\n", ciRequest.Content)
	fmt.Printf("ciRequest.Mask: %q\n", ciRequest.Mask)

	fmt.Printf("%-16s: %08b\n", "Content", ciRequest.Content)
	fmt.Printf("%-16s: %08b\n", "Mask", ciRequest.Mask)

	str := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%08b", ciRequest.Content)))
	fmt.Printf("byte base64: %q\n", str)

	_, _, err = client.Collection("context").Add(ctx, map[string]interface{}{
		"content_bytes":          fmt.Sprintf("%08b", ciRequest.Content),
		"mask_bytes":             fmt.Sprintf("%08b", ciRequest.Mask),
		"content_base64_encoded": base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%08b", ciRequest.Content))),
		"mask_base64_encoded":    base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%08b", ciRequest.Mask))),
		"UnixTime":               time.Now().Unix(),
		"number":                 1608588000123456,
	})
	if err != nil {
		log.Fatalf("Failed adding context: %v", err)
	}

	// [START fs_get_all_users]
	iter := client.Collection("context").Documents(ctx)
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

	query := client.Collection("context").
		Where("content_bytes", ">=", "00000000").
		OrderBy("content_bytes", firestore.Asc).
		Where("content_bytes", "<=", "00000100").
		Select("content_bytes")

	iter = query.Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		fmt.Printf("doc.Ref.ID: %q\n", doc.Ref.ID)

		//	if err != nil {
		//		log.Fatalf("Failed to iterate: %v", err)
		//	}
		//	//fmt.Printf("where from: %q\n\n", doc.Data())
		//	if doc.Data()["to"].(time.Time).Unix() <= to.Unix() {
		//
		//		fmt.Printf("between yesterday and tomorrow: %q\n\n", doc.Data())
		//	} else {
		//
		fmt.Printf("NOT: %q\n\n", doc.Data())
		//}
	}

	from := time.Date(2020, time.December, 21, 22, 0, 0, 0, time.UTC)
	to := time.Date(2020, time.December, 21, 23, 0, 0, 0, time.UTC)

	fmt.Printf("%-16s: %v\n", "from", from.Unix())
	fmt.Printf("%-16s: %v\n", "to", to.Unix())
	fmt.Printf("%-16s: %08b\n", "from", from.Unix())
	fmt.Printf("%-16s: %08b\n", "to", to.Unix())

	//fmt.Printf("%-16s: %08b\n", "Content", ciRequest.Content)

	//10111111110000100

}
