package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/voyagen/17live/client"
	"github.com/voyagen/17live/client/event"
)

func main() {
	cfg := client.Config{
		Username: "yourusername",
		Password: "yourpassword",
		Channels: []int{1234567890}, // Livestream
	}

	c, err := client.NewClient(cfg)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.OnMessage(func(client *client.Client, chatmessage *event.ChatMessage) {
		fmt.Println(fmt.Sprintf("[%s] Message - %s: %s", chatmessage.RoomID, chatmessage.Username, chatmessage.Text))
	})

	c.OnPoke(func(client *client.Client, poke *event.Poke) {
		// Make sure that you are the person that was poked.
		if poke.Receiver.UserID == client.User.UserID {
			fmt.Println(fmt.Sprintf("[%s] Poke: %s", poke.RoomID, poke.Sender.DisplayName))
		}
	})

	c.OnRedEnvelopeInfo(func(client *client.Client, envelope *event.RedEnvelopeInfo) {
		openTime := time.Unix(int64(envelope.StartTime), 0).Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf("[%s] Red envelope - %s", envelope.RoomID, openTime))

	})

	c.OnUserJoined(func(client *client.Client, userJoined *event.UserJoined) {
		fmt.Println(fmt.Sprintf("[%s] User Joined - %s", userJoined.RoomID, userJoined.Username))
	})

	ctx := context.Background()
	err = c.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// keeps the app running
	select {}
}
