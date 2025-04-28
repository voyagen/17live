package client

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client with a message handler
type Client struct {
	conn      *websocket.Conn
	onMessage func(msg DecryptedMessage)
}

// NewClient creates a new WebSocket client
func NewClient() (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial("wss://17media-realtime.ably.io/?key=qvDtFQ.0xBeRA:iYWpd3nD2QHE6Sjm&format=json&heartbeats=true&v=1.1&lib=js-web-1.1.25", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %v", err)
	}

	client := &Client{
		conn: conn,
	}
	return client, nil
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) SetOnMessage(handler func(msg DecryptedMessage)) {
	c.onMessage = handler
}

// Start begins listening for messages and triggers the OnMessage handler
func (c *Client) Start() {
	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			HandleMessage(message, c.onMessage)
		}
	}()
}

// Subscribe sends a subscription message to a specific channel
func (c *Client) Subscribe(channel string) error {
	subscribeMessage := fmt.Sprintf(`{
		"action": 10,
		"channel": "%s"
	}`, channel)

	return c.conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
}

// Close terminates the WebSocket connection
func (c *Client) Close() {
	c.conn.Close()
}
