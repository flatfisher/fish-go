package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://cloud.google.com/feeds/cloud-armor-release-notes.xml")
	items := feed.Items
	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println()
	}
}
