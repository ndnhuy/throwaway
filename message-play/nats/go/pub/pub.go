package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	// Use the env variable if running in the container, otherwise use the default.
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	fmt.Println("Connected to NAT: " + url)

	// Create an unauthenticated connection to NATS.
	nc, _ := nats.Connect(url)

	var wg sync.WaitGroup
	wg.Add(1)
	// Drain is a safe way to to ensure all buffered messages that were published
	// are sent and all buffered messages received on a subscription are processed
	// being closing the connection.
	defer nc.Drain()

	// Publish a couple messages.
	nc.Publish("greet.joe", []byte("hello"))
	nc.Publish("greet.pam", []byte("hello"))

	// One more for good measures..
	nc.Publish("greet.bob", []byte("hello"))
	wg.Wait()
	fmt.Println("Exit")
}
