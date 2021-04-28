package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

const (
	gcpProjectID = "nlsn-connect-data-eng-poc"
	topicID      = "unicorn_topic"
)

func publishtoGcp() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, gcpProjectID)
	if err != nil {
		panic(err)
	}

	//var wg sync.WaitGroup
	//var totalErrors uint64
	t := client.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello, World!! GCP Pub/Sub"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})

	id, err := result.Get(ctx)
	if err != nil {
		fmt.Printf("Get:%v", err)
	}
	fmt.Printf("Published message with custom attributes. msg id: %v\n", id)
}
