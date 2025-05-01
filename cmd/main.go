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
		Channels: []int{1234567890}, // Livestream
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
