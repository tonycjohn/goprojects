package main

import (
	"context"
	"fmt"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
)

const (
	connStr = "Endpoint=sb://unicorn.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xwFlR1hBRjOWlTTFdK8rx6W8d8tiznVscShacZ0R4jg="
)

func sendMessageToQue() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		panic(err)
	}

	q, err := ns.NewQueue("unicorn_data_queue")
	if err != nil {
		panic(err)
	}

	err = q.Send(ctx, servicebus.NewMessageFromString("Hello Azure Service Bus Queue, from Go!!"))
	if err != nil {
		panic(err)
	}

}

func readFromQue() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		panic(err)
	}
	q, err := ns.NewQueue("unicorn_data_queue")
	if err != nil {
		panic(err)
	}

	var printMessage servicebus.HandlerFunc = func(ctx context.Context, msg *servicebus.Message) error {
		fmt.Println(string(msg.Data))
		return msg.Complete(ctx)
	}
	if err := q.Receive(ctx, printMessage); err != nil {
		fmt.Println("FATAL: ", err)
	}
}

func sendMessageToTopic() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		panic(err)
	}

	t, err := ns.NewTopic("unicorn_data_topic")
	if err != nil {
		panic(err)
	}

	err = t.Send(ctx, servicebus.NewMessageFromString("Hello Azure Service Bus Topic!, from Go!!"))
	if err != nil {
		panic(err)
	}
}

func readFromTopic() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		panic(err)
	}

	t, err := ns.NewTopic("unicorn_data_topic")
	if err != nil {
		panic(err)
	}

	subscription, err := t.NewSubscription("unicorn-topic-subscription")
	if err != nil {
		panic(err)
	}
	defer subscription.Close(ctx)

	var printMessage servicebus.HandlerFunc = func(ctx context.Context, msg *servicebus.Message) error {
		fmt.Println(string(msg.Data))
		return msg.Complete(ctx)
	}
	if err = subscription.Receive(ctx, printMessage); err != nil {
		fmt.Println("FATAL: ", err)
	}
}
