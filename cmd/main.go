package main

import (
	"fmt"
	"log"
	"time"

	"github.com/voyagen/17live/client"
)

// EXAMPLE CODE
func main() {
	var clientConfig client.ClientConfig = client.ClientConfig{
		Channels: []int{
			8494770,
		},
	}
	client, err := client.NewClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	client.OnMessage(messageHandler)
	client.Connect()

	for {
		time.Sleep(1 * time.Second)
	}
}

func messageHandler(message client.Message) {
	fmt.Printf("[%s] %s: %s\n", message.Channel, message.Username, message.Content)
}
