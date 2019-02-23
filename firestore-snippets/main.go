package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

type City struct {
	Name       string   `firestore:"name,omitempty"`
	State      string   `firestore:"state,omitempty"`
	Country    string   `firestore:"country,omitempty"`
	Capital    bool     `firestore:"capital,omitempty"`
	Population int64    `firestore:"population,omitempty"`
	Regions    []string `firestore:"regions,omitempty"`
}

func main() {
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJECT")
	if projectID == "" {
		log.Fatalf("Set Firebase project ID via GCLOUD_PROJECT env variable.")
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err)
	}
	defer client.Close()

	if err := prepareQuery(ctx, client); err != nil {
		log.Fatalf("Cannot prepare query docs: %v", err)
	}
	if err := addDocAsMap(ctx, client); err != nil {
		log.Fatalf("Cannot add document as map: %v", err)
	}
	if err := addDocDataTypes(ctx, client); err != nil {
		log.Fatalf("Cannot add document with more data types: %v", err)
	}
	if err := addDocAsEntity(ctx, client); err != nil {
		log.Fatalf("Cannot add document as entity: %v", err)
	}
	if err := addDocWithID(ctx, client); err != nil {
		log.Fatalf("Cannot add doc with id: %v", err)
	}
	if err := addDocWithoutID(ctx, client); err != nil {
		log.Fatalf("Cannot add doc without id: %v", err)
	}
	if err := addDocAfterAutoGeneratedID(ctx, client); err != nil {
		log.Fatalf("Cannot add document after generating ID: %v", err)
	}
	if err := updateDoc(ctx, client); err != nil {
		log.Fatalf("Cannot update doc: %v", err)
	}
	if err := updateDocCreateIfMissing(ctx, client); err != nil {
		log.Fatalf("Cannot update doc, creating if missing: %v", err)
	}
	if err := updateDocMultiple(ctx, client); err != nil {
		log.Fatalf("Cannot update multiple docs: %v", err)
	}
	if err := updateDocNested(ctx, client); err != nil {
		log.Fatalf("Cannot update nested doc: %v", err)
	}
}
