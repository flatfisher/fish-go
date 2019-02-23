package main

import (
	"context"
	"log"
	"time"

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

func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
	doc := make(map[string]interface{})
	doc["stringExample"] = "Hello world!"
	doc["booleanExample"] = true
	doc["numberExample"] = 3.14159265
	doc["dateExample"] = time.Now()
	doc["arrayExample"] = []interface{}{5, true, "hello"}
	doc["nullExample"] = nil
	doc["objectExample"] = map[string]interface{}{
		"a": 5,
		"b": true,
	}

	_, err := client.Collection("data").Doc("one").Set(ctx, doc)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return err
}

func addDocAsEntity(ctx context.Context, client *firestore.Client) error {
	city := City{
		Name:    "Los Angeles",
		Country: "USA",
	}
	_, err := client.Collection("cities").Doc("LA").Set(ctx, city)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return err
}

func addDocWithID(ctx context.Context, client *firestore.Client) error {
	var data = make(map[string]interface{})
	_, err := client.Collection("cities").Doc("new-city-id").Set(ctx, data)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return err
}

func addDocWithoutID(ctx context.Context, client *firestore.Client) error {
	_, _, err := client.Collection("cities").Add(ctx, map[string]interface{}{
		"name":    "Tokyo",
		"country": "Japan",
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return err
}
