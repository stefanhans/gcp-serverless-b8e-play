package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/genproto/googleapis/type/latlng"
	"log"
	"time"
)

type Member struct {
	ID   int64  `firestore:"id"`
	Name string `firestore:"name"`
}

type Example struct {
	Number    int32          `firestore:"number"`
	String    string         `firestore:"string"`
	Geopoint  *latlng.LatLng `firestore:"geopoint"`
	Timestamp time.Time      `firestore:"timestamp"`
	Struct    Member         `firestore:"member"`
}

func getDocument(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	if collection == "" {
		fmt.Printf("no collection set - use %q\n", "setcollection")
		return
	}
	docId := "vR1EPbLoYRdODb3vUKBO"
	doc, err := client.Collection(collection).Doc(docId).Get(ctx)
	if err != nil {

		fmt.Printf("failed to get document: %v\n", err)
	}
	fmt.Println(doc.Data())

	var s Example
	if err := doc.DataTo(&s); err != nil {
		fmt.Printf("failed to convert data: %v\n", err)
	}
	fmt.Println(s)

}

func cloneDocument(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	if collection == "" {
		fmt.Printf("no collection set - use %q\n", "setcollection <collection>")
		return
	}

	docId := "vR1EPbLoYRdODb3vUKBO"
	//if len(arguments) == 0 {
	//	fmt.Printf("no document id set - use %q\n", "cloneDocument <document id>")
	//	return
	//}

	doc, err := client.Collection(collection).Doc(docId).Get(ctx)
	if err != nil {

		fmt.Printf("failed to get document: %v\n", err)
	}
	fmt.Println(doc.Data())

	var data Example
	if err := doc.DataTo(&data); err != nil {
		fmt.Printf("failed to convert data: %v\n", err)
	}
	fmt.Println(data)

	//var member Member
	m, err := doc.DataAt("Struct")
	if err != nil {
		fmt.Printf("failed to get data at %q: %v\n", "Struct", err)
	}
	fmt.Printf("data.struct: %v (%t)\n", m, m)

	dataMap := doc.Data()
	fmt.Printf("dataMap['Struct']: %v\n", dataMap["Struct"])

	member, ok := dataMap["Struct"].(map[string]interface{})
	if ok {
		data.Struct.ID = member["ID"].(int64)
		data.Struct.Name = member["Name"].(string)
	}

	//data.Struct.ID = m["ID"]

	//err = json.Unmarshal([]byte(m), &result)
	//if err != nil {
	//	fmt.Printf("failed to unmarshall 'jsonRequest': %v\n", err)
	//	return
	//}

	data.String = "another string for the clone"
	if data.Geopoint.Latitude == 0.0 {
		fmt.Printf("data.Geopoint.Latitude: %v\n", data.Geopoint.Latitude)
	}
	if data.Geopoint.Longitude == 0.0 {
		fmt.Printf("data.Geopoint.Longitude: %v\n", data.Geopoint.Longitude)
	}

	fmt.Printf("data.Struct.ID: %v\n", data.Struct.ID)

	data.Struct.ID = 111

	jsonRequest, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshall 'jsonRequest': %v\n", err)
		return
	}
	fmt.Printf("jsonRequest: %s\n", jsonRequest)

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal([]byte(string(jsonRequest)), &result)
	if err != nil {
		fmt.Printf("failed to unmarshall 'jsonRequest': %v\n", err)
		return
	}

	// [START fs_add_data_1]
	_, _, err = client.Collection(collection).Add(ctx, result)
	if err != nil {
		log.Fatalf("Failed adding document to collection %q: %v", collection, err)
	}
}
