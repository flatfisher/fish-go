package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
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
