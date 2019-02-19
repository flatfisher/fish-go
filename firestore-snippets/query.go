package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func prepareQuery(ctx context.Context, client *firestore.Client) error {
	cities := []struct {
		id string
		c  City
	}{
		{
			id: "SF",
			c: City{Name: "San Francisco", State: "CA", Country: "USA",
				Capital: false, Population: 860000,
				Regions: []string{"west_coast", "norcal"}},
		},
		{
			id: "LA",
			c: City{Name: "Los Angeles", State: "CA", Country: "USA",
				Capital: false, Population: 3900000,
				Regions: []string{"west_coast", "socal"}},
		},
		{
			id: "DC",
			c: City{Name: "Washington D.C.", Country: "USA",
				Capital: false, Population: 680000,
				Regions: []string{"east_coast"}},
		},
		{
			id: "TOK",
			c: City{Name: "Tokyo", Country: "Japan",
				Capital: true, Population: 9000000,
				Regions: []string{"kanto", "honshu"}},
		},
		{
			id: "BJ",
			c: City{Name: "Beijing", Country: "China",
				Capital: true, Population: 21500000,
				Regions: []string{"jingjinji", "hebei"}},
		},
	}
	for _, c := range cities {
		if _, err := client.Collection("cities").Doc(c.id).Set(ctx, c.c); err != nil {
			return err
		}
	}
	return nil
}
