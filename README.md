# 17Live Go Client

A Go client for interacting with the 17Live platform, enabling developers to build applications that can receive and send messages, handle red envelope information, send pokes, share reactions, and follow or unfollow streamers.

## Installation

To install the 17Live Go client, run the following command:

```bash
go get github.com/voyagen/17live
```

## Usage

Below is an example demonstrating how to use the client to connect to a 17Live stream, listen for events, and interact with the platform:

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

## WebSocket Event Handlers

The client provides methods to register callbacks for handling real-time WebSocket events:

- `OnMessage`: Registers a callback for handling incoming chat message events.
- `OnRedEnvelopeInfo`: Registers a callback for handling red envelope event data.
- `OnPoke`: Registers a callback for handling poke event notifications.
- `OnUserJoined`: Registers a callback for handling user join event notifications.

## API Methods

The client provides the following methods for interacting with 17Live:

- `SendMessage(roomID int, comment string)`: Send a message to a specific room.
- `PokeAll(roomID int)`: Send a poke to all users in a stream.
- `Poke(userID string, roomID int)`: Send a poke to a specific user in a stream.
- `PokeBack(roomID int)`: Send a poke back to a user who poked you.
- `ShareFacebook(roomID int)`: Share a Facebook reaction for a stream.
- `Share17Live(roomID int)`: Share a 17Live-specific reaction for a stream.
- `Like(roomID int)`: Send a like reaction to a stream.
- `Follow(roomID int)`: Follow a streamer.
- `Unfollow(roomID int)`: Unfollow a streamer.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

For questions, bug reports, or feature requests, please open an issue on the GitHub repository.
