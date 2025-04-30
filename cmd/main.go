package main

import (
	"log"
	"time"

	"github.com/voyagen/17live/client"
)

func main() {
	//Create a configuration from the client
	var cfg client.ClientConfig = client.ClientConfig{
		Username: "yournewfriend",
		Password: "Hallo1234",
		Channels: []int{
			26096429,
		},
	}

	// Create a new client
	client, err := client.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Set on message handler

	// Connect to the channels
	client.Connect()
	client.PokeBack(26096429)
	// This is just to keep the application running
	for {
		time.Sleep(1 * time.Second)
	}
}
