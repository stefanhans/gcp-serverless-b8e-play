package functions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestE2eStore(t *testing.T) {

	testCases := map[string]struct {
		jsonSend     []byte
		expectedName string
	}{
		"Alice": {
			jsonSend: []byte(`{
            "DocId": "123",
            "Name": "Id Test",
            "Type": "Testuser",
            "Status": "",
            "Description": "just a test description"
          }`),
			expectedName: "Alice",
		},
	}

	b := byte(123)

	fmt.Printf("%-16s: %x\n", "purpose", b)

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			resp, err := http.Post("https://europe-west3-serverless-devops-play.cloudfunctions.net/add-user", "application/json", bytes.NewBuffer(tc.jsonSend))
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Could not read response body: %v", err)
			}
			t.Logf("Stored successfully: %s", body)
		})
	}
}
