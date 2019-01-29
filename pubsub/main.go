package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	projectID := ""

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	topicName := "slack-to-googlephotos"

	topic := client.Topic(topicName)

	res := topic.Publish(ctx, &pubsub.Message{Data: []byte("payload")})

	fmt.Printf("Topic %v created.\n", res)
	topic.Stop()
}
