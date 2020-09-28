package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/commandline/types"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func GetMasterData(arguments []string) {

	_ = arguments

	if client == nil {
		fmt.Printf("not connected - use %q\n", "connect")
		return
	}

	//
	// users
	//
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
		fmt.Printf("failed to marshal 'userRequest': %v\n", err)
		return
	}
	fmt.Printf("userRequest: %s\n", jsonUsers)

	usersString := fmt.Sprintf("%q: %s", "Users", jsonUsers)

	//
	// vehicles
	//
	collection = "vehicles"

	vehicles := make([]types.Vehicle, 0)

	iter = client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())

		var data types.Vehicle
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}
		vehicles = append(vehicles, data)
	}
	fmt.Printf("vehicles: \n%v\n", vehicles)

	jsonVehicles, err := json.MarshalIndent(vehicles, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal 'vehicleRequest': %v\n", err)
		return
	}
	fmt.Printf("vehicleRequest: %s\n", jsonVehicles)

	vehiclesString := fmt.Sprintf("%q: %s", "Vehicles", jsonVehicles)

	//
	// bookings
	//
	collection = "bookings"

	bookings := make([]types.Booking, 0)

	iter = client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())

		var data types.Booking
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to convert data: %v\n", err)
		}
		bookings = append(bookings, data)
	}
	fmt.Printf("bookings: \n%v\n", bookings)

	jsonBookings, err := json.MarshalIndent(bookings, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal 'bookingRequest': %v\n", err)
		return
	}
	fmt.Printf("bookingRequest: %s\n", jsonBookings)

	bookingsString := fmt.Sprintf("%q: %s", "Bookings", jsonBookings)

	masterDataString := fmt.Sprintf("{ %s,\n%s,\n%s, \n%q: %q, \n%q: %q, \n%q: %q, \n%q: %q }",
		usersString, vehiclesString, bookingsString,
		"From", time.Now(), "To", time.Now(), "Status", "no-cache", "StatusTime", time.Now())

	out := fmt.Sprintf("{ \n%q: %s }", "MasterData", masterDataString)
	fmt.Printf("------\n%s\n", out)

}
