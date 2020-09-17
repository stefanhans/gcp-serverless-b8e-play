package main

import (
	"fmt"
	"time"

	"github.com/stefanhans/gcp-serverless-b8e-play/Play/booking/types"
)

func main() {

	user := Member{ID: 1,
		Name: "Alice",
	}

	item := Item{ID: 1,
		Name: "Tesla Deluxe",
		Type: CAR,
	}

	bookingRequest := Booking{
		ID:         1,
		User:       user,
		Share:      item,
		From:       time.Now(),
		To:         time.Now().Add(time.Hour),
		Status:     REQUESTED,
		StatusTime: time.Now(),
	}

	fmt.Printf("bookingRequest: %s\n", bookingRequest)
}
