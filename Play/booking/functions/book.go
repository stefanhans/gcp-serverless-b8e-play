package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stefanhans/gcp-serverless-b8e-play/Play/booking/types"
)

func Book(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("bookingRequest: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var bookingRequest types.Booking
	err = json.Unmarshal(bytes, &bookingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshall JSON input: %s", err), http.StatusInternalServerError)
		return
	}

	// Fake the booking procedure as successful
	bookingRequest.Status = types.CONFIRMED

	// Marshal booking
	bookingJson, err := json.Marshal(bookingRequest)
	if err != nil {
		fmt.Printf("failed to marshall 'bookingRequest': %v\n", err)
		return
	}
	fmt.Printf("bookingReply: %s\n", bookingJson)

	// Response registration
	_, err = fmt.Fprint(w, string(bookingJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}
