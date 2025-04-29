package client

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

// NewClient creates a new WebSocket client
// TODO: WORK IN PROGRESS (will merge with api pkg)
func NewClient(config ClientConfig) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial("wss://17media-realtime.ably.io/?key=qvDtFQ.0xBeRA:iYWpd3nD2QHE6Sjm&format=json&heartbeats=true&v=1.1&lib=js-web-1.1.25", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %v", err)
	}

	client := &Client{
		conn:     conn,
		channels: config.Channels,
	}

	return client, nil
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnMessage(handler func(message Message)) {
	c.onMessage = handler
}

// Start begins listening for messages and triggers the OnMessage handler
func (c *Client) Connect() {
	for _, channel := range c.channels {
		if err := c.subscribe(channel); err != nil {
			log.Printf("failed to subscribe to channel %d: %v", channel, err)
		} else {
			log.Printf("Connected to: %d", channel)
		}

	}

	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			HandleOnMessage(message, c.onMessage)
		}
	}()
}

// Close terminates the WebSocket connection
func (c *Client) Disconnect() {
	c.conn.Close()
}

// Subscribe sends a subscription message to a specific channel
func (c *Client) subscribe(channel int) error {
	subscribeMessage := fmt.Sprintf(`{
		"action": 10,
		"channel": "%s"
	}`, strconv.Itoa(channel))

	return c.conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
}
