package main

import (
	"context"
	"log"

	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
)

func connect(arguments []string) {

	// Sets your Google Cloud Platform project ID.
	projectID := "serverless-devops-play"

	// Get a Firestore client.
	ctx = context.Background()

	sa := option.WithCredentialsFile("/Users/stefan/.secret/serverless-devops-play-firestore-play.json")
	client, err = firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func disconnect(arguments []string) {
	// Close client
	client.Close()
}
