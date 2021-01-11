package functions

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/functions/types"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"time"
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

	// Todo: Check booking for conflicts and correctness

	// Check requesting user exist
	qUsers := client.Collection("users").
		Where("DocId", "==", bookingRequest.User).
		Limit(1)
	iterUsers := qUsers.Documents(ctx)

	userCheck, err := iterUsers.Next()
	if err != nil {
		http.Error(w, fmt.Sprintf("unknown user: %s", bookingRequest.User), http.StatusInternalServerError)
		return
	}
	_ = userCheck

	// Check requested vehicle exist
	qVehicles := client.Collection("vehicles").
		Where("DocId", "==", bookingRequest.Vehicle).
		Limit(1)
	iterVehicles := qVehicles.Documents(ctx)

	vehicleCheck, err := iterVehicles.Next()
	if err != nil {
		http.Error(w, fmt.Sprintf("unknown vehicle: %s", bookingRequest.Vehicle), http.StatusInternalServerError)
		return
	}
	_ = vehicleCheck

	// Check for booking conflicts
	qBookings := client.Collection("bookings").
		Where("vehicle", "==", bookingRequest.Vehicle)
	iterBookings := qBookings.Documents(ctx)

	for {
		doc, err := iterBookings.Next()
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

		// Exist booking with "from" within requested range
		if data.From.Unix() >= bookingRequest.From.Unix() && data.From.Unix() <= bookingRequest.To.Unix() {
			http.Error(w, fmt.Sprintf("booking conflict (1) found for vehicle %s", bookingRequest.Vehicle), http.StatusInternalServerError)
			return
		}

		// Exist booking with "to" within requested range
		if data.To.Unix() >= bookingRequest.From.Unix() && data.To.Unix() <= bookingRequest.To.Unix() {
			http.Error(w, fmt.Sprintf("booking conflict (2) found for vehicle %s", bookingRequest.Vehicle), http.StatusInternalServerError)
			return
		}

		// Exist booking with "from" before and "to" after requested range
		if data.From.Unix() <= bookingRequest.From.Unix() && data.To.Unix() >= bookingRequest.To.Unix() {
			http.Error(w, fmt.Sprintf("booking conflict (3) found for vehicle %s", bookingRequest.Vehicle), http.StatusInternalServerError)
			return
		}
	}

	doc := client.Doc("bookings/" + fmt.Sprintf("%v", bookingRequest.FromToQuarters()))

	// Set the document id
	bookingRequest.DocId = doc.ID

	// Set the registration status and (booking) time
	bookingRequest.Status = "registered from"
	bookingRequest.StatusTime = time.Now()

	writeResult, err := doc.Create(ctx, bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create bookings/%s: %s", bookingRequest.DocId, err), http.StatusInternalServerError)
		return
	}
	_ = writeResult

	doc = client.Doc("bookings/" + fmt.Sprintf("%v", bookingRequest.ToToQuarters()))

	// Set the document id
	bookingRequest.DocId = doc.ID

	// Set the registration status and (booking) time
	bookingRequest.Status = "registered to"
	bookingRequest.StatusTime = time.Now()

	writeResult, err = doc.Create(ctx, bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create bookings/%s: %s", bookingRequest.DocId, err), http.StatusInternalServerError)
		return
	}
	_ = writeResult

	//// Check for predefined booking resp. document id
	//if bookingRequest.DocId == "" {
	//
	//	doc := client.Collection("bookings").NewDoc()
	//
	//	// Set the document id
	//	bookingRequest.DocId = doc.ID
	//
	//	// Set the registration status and (booking) time
	//	bookingRequest.Status = "registered"
	//	bookingRequest.StatusTime = time.Now()
	//
	//	// Register the booking
	//	wr, err := doc.Create(ctx, bookingRequest)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("failed to create document: %s", err), http.StatusInternalServerError)
	//		return
	//	}
	//	_ = wr
	//	//fmt.Println(wr.UpdateTime)
	//
	//} else {
	//
	//	doc := client.Doc("bookings/" + bookingRequest.DocId)
	//	writeResult, err := doc.Create(ctx, bookingRequest)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("failed to create bookings/%s: %s", bookingRequest.DocId, err), http.StatusInternalServerError)
	//		return
	//	}
	//	_ = writeResult
	//
	//}

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

func GetBookingById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("GetBookingById Request: %s\n", body)

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

	// "49KvH6m2Y7flEegBdPv4"
	q := client.Collection("bookings").
		Where("DocId", "==", bookingRequest.DocId).
		Limit(1)
	iter := q.Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		http.Error(w, fmt.Sprintf("no document found: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(doc.Data())
	booking := doc.Data()

	fmt.Printf("booking: \n%v\n", booking)

	jsonBooking, err := json.MarshalIndent(booking, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonBooking': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("jsonBooking: %s\n", jsonBooking)

	// Response
	_, err = fmt.Fprint(w, string(jsonBooking))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func DeleteBookingById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("DeleteBookingById Request: %s\n", body)

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

	deleteResult, err := client.Batch().
		Delete(client.Doc("bookings/"+bookingRequest.DocId), firestore.Exists).
		Commit(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete booking %q: %s", bookingRequest.DocId, err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("C: \n%v\n", deleteResult)
	bookingRequest.Status = "deleteResult: \n"
	for i, s := range deleteResult {
		bookingRequest.Status += fmt.Sprintf("%v: %q\n", i, s)
	}

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

func GetBookingsByRange(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("GetBookingsByRange Request: %s\n", body)

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

	bookings := make([]types.Booking, 0)

	// "49KvH6m2Y7flEegBdPv4"
	q := client.Collection("bookings").
		Where("To", ">=", bookingRequest.From)

	iter := q.Documents(ctx)

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

		// Filter the result by "from" as endtime
		if data.From.Unix() <= bookingRequest.To.Unix() {
			bookings = append(bookings, data)
		}
	}

	jsonBookings, err := json.MarshalIndent(bookings, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonBookings': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("jsonBookings: %s\n", jsonBookings)

	// Response
	_, err = fmt.Fprintf(w, "{ \n    %q: %s\n}\n", "Bookings", string(jsonBookings))
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
	fmt.Printf("GetBookings Request: %s\n", body)

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

		var booking types.Booking
		if err := doc.DataTo(&booking); err != nil {
			http.Error(w, fmt.Sprintf("failed to convert data: %s", err), http.StatusInternalServerError)
			return
		}
		booking.Status = fmt.Sprintf("booking.FromToQuarters(): %v", booking.FromToQuarters())
		bookings = append(bookings, booking)
	}
	fmt.Printf("bookings: \n%v\n", bookings)

	jsonBookings, err := json.MarshalIndent(bookings, "    ", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonBookings': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("bookingsRequest: %s\n", jsonBookings)

	// Response
	_, err = fmt.Fprintf(w, "{ \n    %q: %s\n}\n", "Bookings", string(jsonBookings))
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
