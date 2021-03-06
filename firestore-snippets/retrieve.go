package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func prepareRetrieve(ctx context.Context, client *firestore.Client) error {
	// [START fs_retrieve_create_examples]
	cities := []struct {
		id string
		c  City
	}{
		{id: "SF", c: City{Name: "San Francisco", State: "CA", Country: "USA", Capital: false, Population: 860000}},
		{id: "LA", c: City{Name: "Los Angeles", State: "CA", Country: "USA", Capital: false, Population: 3900000}},
		{id: "DC", c: City{Name: "Washington D.C.", Country: "USA", Capital: true, Population: 680000}},
		{id: "TOK", c: City{Name: "Tokyo", Country: "Japan", Capital: true, Population: 9000000}},
		{id: "BJ", c: City{Name: "Beijing", Country: "China", Capital: true, Population: 21500000}},
	}
	for _, c := range cities {
		_, err := client.Collection("cities").Doc(c.id).Set(ctx, c.c)
		if err != nil {
			return err
		}
	}
	return nil
}

func docAsMap(ctx context.Context, client *firestore.Client) (map[string]interface{}, error) {
	dsnap, err := client.Collection("cities").Doc("SF").Get(ctx)
	if err != nil {
		return nil, err
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	return m, nil
}

func docAsEntity(ctx context.Context, client *firestore.Client) (*City, error) {
	dsnap, err := client.Collection("cities").Doc("BJ").Get(ctx)
	if err != nil {
		return nil, err
	}
	var c City
	dsnap.DataTo(&c)
	fmt.Printf("Document data: %#v\n", c)
	return &c, nil
}

func multipleDocs(ctx context.Context, client *firestore.Client) error {
	fmt.Println("All capital cities:")
	iter := client.Collection("cities").Where("capital", "==", true).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

func allDocs(ctx context.Context, client *firestore.Client) error {
	fmt.Println("All cities:")
	iter := client.Collection("cities").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

func getCollections(ctx context.Context, client *firestore.Client) error {
	iter := client.Collection("cities").Doc("SF").Collections(ctx)
	for {
		collRef, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Found collection with id: %s\n", collRef.ID)
	}
	return nil
}
