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
		Channels: []int{1234567890},
	}

	c, err := client.NewClient(cfg)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.OnMessage(func(client *client.Client, chatmessage *event.ChatMessage) {
		fmt.Println("Message:", chatmessage.Username, chatmessage.Text)
	})

	c.OnRedEnvelopeInfo(func(client *client.Client, envelope *event.RedEnvelopeInfo) {
		openTime := time.Unix(int64(envelope.StartTime), 0).Format("2006-01-02 15:04:05")
		endTime := time.Unix(int64(envelope.StartTime), 0).Format("2006-01-02 15:04:05")
		fmt.Println("Red Envelope:", openTime, endTime)
	})

	c.Connect()
}
