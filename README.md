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
