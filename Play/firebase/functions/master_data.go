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

func GetMasterData(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("AddUser Request: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var userRequest types.User
	err = json.Unmarshal(bytes, &userRequest)
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

	//
	// users
	//
	users := make([]types.User, 0)

	iter := client.Collection("users").Documents(ctx)
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

		var data types.User
		if err := doc.DataTo(&data); err != nil {
			http.Error(w, fmt.Sprintf("failed to convert data: %s", err), http.StatusInternalServerError)
			return
		}
		users = append(users, data)
	}
	fmt.Printf("users: \n%v\n", users)

	jsonUsers, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonUsers': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("usersRequest: %s\n", jsonUsers)

	usersString := fmt.Sprintf("%q: %s", "Users", jsonUsers)

	//
	// vehicles
	//
	vehicles := make([]types.Vehicle, 0)

	iter = client.Collection("vehicles").Documents(ctx)
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

		var data types.Vehicle
		if err := doc.DataTo(&data); err != nil {
			http.Error(w, fmt.Sprintf("failed to convert data: %s", err), http.StatusInternalServerError)
			return
		}
		vehicles = append(vehicles, data)
	}
	fmt.Printf("vehicles: \n%v\n", vehicles)

	jsonVehicles, err := json.MarshalIndent(vehicles, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonVehicles': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("vehiclesRequest: %s\n", jsonVehicles)

	vehiclesString := fmt.Sprintf("%q: %s", "Vehicles", jsonVehicles)

	//
	// bookings
	//
	bookings := make([]types.Booking, 0)

	iter = client.Collection("bookings").Documents(ctx)
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
	fmt.Printf("bookingsRequest: %s\n", jsonBookings)

	bookingsString := fmt.Sprintf("%q: %s", "Bookings", jsonBookings)

	masterDataString := fmt.Sprintf("{ %s,\n%s,\n%s, \n%q: %q, \n%q: %q, \n%q: %q, \n%q: %q }",
		usersString, vehiclesString, bookingsString,
		"From", time.Now(), "To", time.Now(), "Status", "no-cache", "StatusTime", time.Now())

	// Response
	_, err = fmt.Fprintf(w, "{ \n%q: %s }\n", "MasterData", masterDataString)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}

}
