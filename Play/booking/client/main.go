package main

import (
	"encoding/json"
	"fmt"
	"github.com/stefanhans/gcp-serverless-b8e-play/Play/booking/types"
	"io/ioutil"
	"net/http"
	"strings"
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

	jsonRequest, err := json.Marshal(bookingRequest)
	if err != nil {
		fmt.Printf("failed to marshall 'bookingRequest': %v\n", err)
		return
	}
	//fmt.Printf("bookingRequest: %s\n", json)

	// Send request to service
	res, err := http.Post("http://localhost:8080/",
		"application/json",
		strings.NewReader(fmt.Sprintf("%s", jsonRequest)))
	if err != nil {
		fmt.Printf("failed to POST JSON: %v\n", err)
		return
	}

	// Read response body in JSON
	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		fmt.Printf("failed to read response: %v\n", err)
		return
	}
	//fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var bookingReply types.Booking
	err = json.Unmarshal(bytes, &bookingReply)
	if err != nil {
		fmt.Printf("cannot unmarshall JSON input: %s", err)
		return
	}

	// Prints the reply
	fmt.Printf("Booking of %s %q in status %q: %s\n", bookingReply.Share.Type, bookingReply.Share.Name, bookingReply.Status, body)
}
