package functions

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"

	"github.com/stefanhans/gcp-serverless-b8e-play/Play/firebase/functions/types"
	"io/ioutil"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	fmt.Printf("GetUserById Request: %s\n", body)

	if err != nil {
		fmt.Printf("failed to read request: %s", err)
		_, err = fmt.Fprint(w,
			types.CreateResponse("error",
				fmt.Sprintf("failed to read request: %s", err),
				body))
		return
	}

	// Unmarshal request body
	bytes := []byte(string(body))
	var userRequest types.User

	err = json.Unmarshal(bytes, &userRequest)
	if err != nil {
		fmt.Printf("cannot unmarshall JSON input: %s", err)
		_, err = fmt.Fprint(w,
			types.CreateResponse("error",
				fmt.Sprintf("cannot unmarshall JSON input: %s", err),
				body))
		return
	}

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	//sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("failed to create firestore client: %s", err)
		_, err = fmt.Fprint(w,
			types.CreateResponse("error",
				fmt.Sprintf("failed to create firestore client: %s", err),
				body))
		return
	}

	// "49KvH6m2Y7flEegBdPv4"
	q := client.Collection("users").
		Where("DocId", "==", userRequest.DocId).
		Limit(1)
	iter := q.Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		fmt.Printf("no document found: %s", err)
		_, err = fmt.Fprint(w,
			types.CreateResponse("error",
				fmt.Sprintf("no document found: %s", err),
				body))
		return
	}
	user := doc.Data()

	jsonUser, err := json.MarshalIndent(user, "    ", "    ")
	if err != nil {
		fmt.Printf("failed to marshal firestore data: %s", err)
		_, err = fmt.Fprint(w,
			types.CreateResponse("error",
				fmt.Sprintf("failed to marshal firestore data: %s", err),
				[]byte(fmt.Sprintf("\"%v\"", user))))
		return
	}

	//var userObject types.User
	//err = json.Unmarshal(jsonUser, userObject)
	//
	//userObject.Description

	// Response
	fmt.Print(types.CreateResponse("success", "", jsonUser))
	_, err = fmt.Fprint(w, types.CreateResponse("success", "", jsonUser))
	if err != nil {
		fmt.Printf("failed to write response: %s", err)
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}
