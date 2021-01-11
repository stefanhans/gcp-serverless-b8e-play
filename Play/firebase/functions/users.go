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

	// Check for predefined user resp. document id
	if userRequest.DocId == "" {

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

	} else {

		doc := client.Doc("users/" + userRequest.DocId)
		writeResult, err := doc.Create(ctx, userRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to create users/%s: %s", userRequest.DocId, err), http.StatusInternalServerError)
			return
		}
		_ = writeResult

	}

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

func GetUserById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("GetUserById Request: %s\n", body)

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

	// "49KvH6m2Y7flEegBdPv4"
	q := client.Collection("users").
		Where("DocId", "==", userRequest.DocId).
		Limit(1)
	iter := q.Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		http.Error(w, fmt.Sprintf("no document found: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(doc.Data())
	user := doc.Data()

	fmt.Printf("user: \n%v\n", user)

	jsonUser, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonUser': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("userRequest: %s\n", jsonUser)

	// Response
	_, err = fmt.Fprint(w, string(jsonUser))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("DeleteUserById Request: %s\n", body)

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

	deleteResult, err := client.Batch().
		Delete(client.Doc("users/"+userRequest.DocId), firestore.Exists).
		Commit(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete user %q: %s", userRequest.DocId, err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("C: \n%v\n", deleteResult)
	userRequest.Description = "deleteResult: \n"
	for i, s := range deleteResult {
		userRequest.Description += fmt.Sprintf("%v: %q\n", i, s)
	}

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
		_, err = fmt.Fprintf(w, "{%q: %q}", "result", "cleared")
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
	fmt.Printf("GetUsers Request: %s\n", body)

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

	jsonUsers, err := json.MarshalIndent(users, "    ", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal data 'jsonUsers': %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("userRequest: %s\n", jsonUsers)

	// Response
	_, err = fmt.Fprintf(w, "{ \n    %q: %s\n}\n", "Users", string(jsonUsers))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
}
