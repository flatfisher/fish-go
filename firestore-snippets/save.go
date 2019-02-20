package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func addDocAsMap(ctx context.Context, client *firestore.Client) error {
	// [START fs_add_simple_doc_as_map]
	_, err := client.Collection("cities").Doc("LA").Set(ctx, map[string]interface{}{
		"name":    "Los Angeles",
		"state":   "CA",
		"country": "USA",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_simple_doc_as_map]
	return err
}
