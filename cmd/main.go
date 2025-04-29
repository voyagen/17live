package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/voyagen/17live/client"
)

// EXAMPLE CODE
func main() {
	var clientConfig client.ClientConfig = client.ClientConfig{
		Username: "username",
		Password: "password",
		Channels: []int{
			28743507,
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

func messageHandler(c *client.Client, message client.Message) {
	fmt.Printf("[%s] %s: %s\n", message.Channel, message.Username, message.Content)

	channelID, err := strconv.Atoi(message.Channel)
	if err != nil {
		log.Printf("Invalid channel ID: %s\n", message.Channel)
		return
	}

	text := fmt.Sprintf("@%v Hello!", message.Username)
	c.SendMessage(channelID, text)
}
