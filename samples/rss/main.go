package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://cloud.google.com/feeds/stackdriver-release-notes.xml")
	items := feed.Items

	for _, item := range items {
		fmt.Println(item.Content)
		fmt.Println()
	}
}
