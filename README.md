# 17Live Go Client

A Go client for interacting with the 17Live API, providing functionality to receive messages, send messages, send pokes, share reactions, and follow and unfollow streamers.

## Installation

To install the 17Live Go client, use:

```bash
go get github.com/voyagen/17live
```

## Usage

Below is an example of how to use the client to send a message:

```go
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/voyagen/17live/client"
)

func main() {
    //Create a configuration from the client
	var cfg client.ClientConfig = client.ClientConfig{
		Username: "username",
		Password: "password",
		Channels: []int{
			123456789,
		},
	}

    // Create a new client
	client, err := client.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

    // Set on message handler
	client.OnMessage(messageHandler)

    // Connect to the channels
	client.Connect()

    // This is just to keep the application running
	for {
		time.Sleep(1 * time.Second)
	}
}

func messageHandler(c *client.Client, message client.Message) {
    // print incomming message
	fmt.Printf("[%s] %s: %s\n", message.Channel, message.Username, message.Content)

	channelID, err := strconv.Atoi(message.Channel)
	if err != nil {
		log.Printf("Invalid channel ID: %s\n", message.Channel)
		return
	}

    // sending a message to back with hello
	text := fmt.Sprintf("@%v Hello!", message.Username)
	c.SendMessage(channelID, text)
}

```

## Features

- Receive messages
- Send messages
- Send poke requests
- Share reactions (Facebook, 17Live)
- Follow and Unfollow streamers

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
