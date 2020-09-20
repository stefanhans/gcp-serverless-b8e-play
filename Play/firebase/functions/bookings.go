package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/functions/types"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
)

func AddBooking(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("AddBooking Request: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var bookingRequest types.Booking
	err = json.Unmarshal(bytes, &bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshall JSON input: %s", err), http.StatusInternalServerError)
		return
	}

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	doc := client.Collection("bookings").NewDoc()

	// Set the document id
	bookingRequest.DocId = doc.ID

	// Set the registration status
	bookingRequest.Status = "registered"

	wr, err := doc.Create(ctx, bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create document: %s", err), http.StatusInternalServerError)
		return
	}
	_ = wr
	//fmt.Println(wr.UpdateTime)

	// Marshal booking
	jsonBooking, err := json.Marshal(bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal 'bookingRequest': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("bookingReply: %s\n", jsonBooking)

	// Response
	_, err = fmt.Fprint(w, string(jsonBooking))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func ClearBookings(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClearBookings Request: %s\n", body)

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	ref := client.Collection("bookings")
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
				http.Error(w, fmt.Sprintf("failed next iteration: %s", err), http.StatusInternalServerError)
				return
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			fmt.Printf("all %s deleted\n", "bookings")
			return
		}

		result, err := batch.Commit(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to commit batch for deletion: %s", err), http.StatusInternalServerError)
			return
		}
		_ = result
		//fmt.Printf("Commit Result: %v\n", result)

		// Response
		_, err = fmt.Fprintf(w, "{%q: %q}", "result", "cleared")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
		}
	}
}

func GetBookings(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClearBookings Request: %s\n", body)

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %s", err), http.StatusInternalServerError)
		return
	}

	bookings := make([]types.Booking, 0)

	iter := client.Collection("bookings").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("failed next iteration: %s", err), http.StatusInternalServerError)
			return
		}
		fmt.Println(doc.Data())

		var data types.Booking
		if err := doc.DataTo(&data); err != nil {
			http.Error(w, fmt.Sprintf("failed to convert data: %s", err), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, data)
	}
	fmt.Printf("bookings: \n%v\n", bookings)

	jsonBookings, err := json.MarshalIndent(bookings, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonBookings': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("bookingRequest: %s\n", jsonBookings)

	// Response
	_, err = fmt.Fprint(w, string(jsonBookings))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

//func StoreUsers(arguments []string) {
//
//	_ = arguments
//
//	if client == nil {
//		fmt.Printf("not connected - use %q\n", "connect")
//		return
//	}
//
//	jsonUsers, err := ioutil.ReadFile("data/bookings.json")
//	if err != nil {
//		fmt.Printf("could not read file %q: %v\n", "data/bookings.json", err)
//		return
//	}
//
//	fmt.Printf(string(jsonUsers))
//
//	// Declared an empty map interface
//	var bookings []types.User
//
//	// Unmarshal or Decode the JSON to the interface.
//	err = json.Unmarshal([]byte(jsonUsers), &bookings)
//	if err != nil {
//		fmt.Printf("failed to unmarshal 'bookings': %v\n", err)
//	}
//
//	fmt.Printf("bookings: %v\n", bookings)
//
//	for i, booking := range bookings {
//
//		doc := client.Collection("bookings").NewDoc()
//		booking.DocId = doc.ID
//
//		wr, err := doc.Create(ctx, booking)
//		if err != nil {
//			fmt.Printf("failed to create document #%v: %v\n", i, err)
//		}
//		fmt.Println(wr.UpdateTime)
//	}
//}
