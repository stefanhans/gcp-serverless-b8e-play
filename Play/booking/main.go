package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/booking/types"
	"time"
)

func main() {

	user := types.Member{ID: 1,
		Name: "Alice",
	}

	item := types.Item{ID: 1,
		Name: "Tesla Deluxe",
		Type: types.CAR,
	}

	bookingRequest := types.Booking{
		ID:         1,
		User:       user,
		Share:      item,
		From:       time.Now(),
		To:         time.Now().Add(time.Hour),
		Status:     types.REQUESTED,
		StatusTime: time.Now(),
	}

	jsonRequest, err := json.MarshalIndent(bookingRequest, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshall 'bookingRequest': %v\n", err)
		return
	}

	fmt.Printf("bookingRequest: %s\n", jsonRequest)
}
