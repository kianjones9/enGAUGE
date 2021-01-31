package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func dashboard(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//ctx := context.Background()

	fmt.Fprintf(w, "Welcome to the Producer Dashboard!")

	//client, err := pubsub.NewClient(ctx, "ardent-strength-303319")
	if err != nil {
		log.Fatal(err)
	}

	//answers := sub(ctx, client, "answers-sub")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		/*
			answers.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
				conn.WriteMessage(messageType, p)
				message.Ack()
			})
		*/

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

	/*
		questions := pub(ctx, client, "questions-pub")

		questions.Publish(ctx, &pubsub.Message{
			Data: []byte("Test Question?"),
		})
	*/

}

func sub(ctx context.Context, client *pubsub.Client, topic string) *pubsub.Subscription {

	subscription := client.Subscription(topic)

	return subscription
}

func pub(ctx context.Context, client *pubsub.Client, topic string) *pubsub.Topic {

	publish := client.Topic(topic)

	return publish

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to enGAUGEment!\n Select a show below!")
}

func handler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/dashboard", dashboard)
}

func main() {

	handler()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
