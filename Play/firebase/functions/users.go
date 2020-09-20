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

func AddUser(w http.ResponseWriter, r *http.Request) {

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

	doc := client.Collection("users").NewDoc()

	// Set the document id
	userRequest.DocId = doc.ID

	// Set the registration status
	userRequest.Status = "registered"

	wr, err := doc.Create(ctx, userRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create document: %s", err), http.StatusInternalServerError)
		return
	}
	_ = wr
	//fmt.Println(wr.UpdateTime)

	// Marshal user
	jsonUser, err := json.Marshal(userRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal 'userRequest': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("userReply: %s\n", jsonUser)

	// Response
	_, err = fmt.Fprint(w, string(jsonUser))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func ClearUsers(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClearUsers Request: %s\n", body)

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

	ref := client.Collection("users")
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
			fmt.Printf("all %s deleted\n", "users")
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
		_, err = fmt.Fprint(w, "{%q: %q}", "result", "cleared")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
		}
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClearUsers Request: %s\n", body)

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
	fmt.Printf("userRequest: %s\n", jsonUsers)

	// Response
	_, err = fmt.Fprint(w, string(jsonUsers))
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
//	jsonUsers, err := ioutil.ReadFile("data/users.json")
//	if err != nil {
//		fmt.Printf("could not read file %q: %v\n", "data/users.json", err)
//		return
//	}
//
//	fmt.Printf(string(jsonUsers))
//
//	// Declared an empty map interface
//	var users []types.User
//
//	// Unmarshal or Decode the JSON to the interface.
//	err = json.Unmarshal([]byte(jsonUsers), &users)
//	if err != nil {
//		fmt.Printf("failed to unmarshal 'users': %v\n", err)
//	}
//
//	fmt.Printf("users: %v\n", users)
//
//	for i, user := range users {
//
//		doc := client.Collection("users").NewDoc()
//		user.DocId = doc.ID
//
//		wr, err := doc.Create(ctx, user)
//		if err != nil {
//			fmt.Printf("failed to create document #%v: %v\n", i, err)
//		}
//		fmt.Println(wr.UpdateTime)
//	}
//}
