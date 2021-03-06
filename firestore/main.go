package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := ""

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	// [START fs_add_data_1]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	// [END fs_add_data_1]

	// [START fs_get_all_users]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	// [END fs_get_all_users]
}
