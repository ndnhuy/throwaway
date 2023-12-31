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

	// Create an unauthenticated connection to NATS.
	nc, _ := nats.Connect(url)

	fmt.Println("Connected to NAT: " + url)
	// Drain is a safe way to to ensure all buffered messages that were published
	// are sent and all buffered messages received on a subscription are processed
	// being closing the connection.
	defer nc.Drain()

	var wg sync.WaitGroup

	// Let's create a subscription on the greet.* wildcard.
	wg.Add(3)
	_, _ = nc.Subscribe("greet.*", func(msg *nats.Msg) {
		fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
		wg.Done()
	})

	wg.Wait()
	fmt.Println("Exit")
}
