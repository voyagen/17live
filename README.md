# 17Live Go Client

A Go client for interacting with the 17Live API, providing functionality to receive messages, send messages, receive red envelope info, send pokes, share reactions, and follow and unfollow streamers.

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
	"time"

	"github.com/voyagen/17live/client"
	"github.com/voyagen/17live/internal/event"
)

func main() {
	cfg := client.Config{
		Username: "yourusername",
		Password: "yourpassword",
		Channels: []int{1234567890}, // Livestream ID can be found in the url -> https://17.live/en/live/1234567890
	}

	c, err := client.NewClient(cfg)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.OnMessage(func(client *client.Client, chatmessage *event.ChatMessage) {
		fmt.Println("Message:", chatmessage.Username, chatmessage.Text)
	})

	c.OnPoke(func(client *client.Client, poke *event.Poke) {
		if poke.Receiver.UserID == client.User.UserID {
			fmt.Println(fmt.Sprintf("You are poked by %s", poke.Sender.DisplayName))
		}
	})

	c.OnRedEnvelopeInfo(func(client *client.Client, envelope *event.RedEnvelopeInfo) {
		openTime := time.Unix(int64(envelope.StartTime), 0).Format("2006-01-02 15:04:05")
		fmt.Println(openTime)
	})

	c.OnUserJoined(func(client *client.Client, userJoined *event.UserJoined) {
		fmt.Println("User Joined:", userJoined.RoomID, userJoined.Username)
	})

	c.Connect()
}


```

## Features

- Receive messages
- Receive Red Envelope Information
- Send messages
- Send poke requests
- Share reactions (Facebook, 17Live)
- Follow and Unfollow streamers

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
