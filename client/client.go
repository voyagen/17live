package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// NewClient creates a new WebSocket client
func NewClient(config ClientConfig) (*Client, error) {
	if config.Username == "" {
		return nil, errors.New("config username cannot be empty")
	}
	if config.Password == "" {
		return nil, errors.New("config password cannot be empty")
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://17media-realtime.ably.io/?key=qvDtFQ.0xBeRA:iYWpd3nD2QHE6Sjm&format=json&heartbeats=true&v=1.1&lib=js-web-1.1.25", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %v", err)
	}

	restyclient := resty.New()
	restyclient.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"devicetype":   "WEB",
		"language":     "GLOBAL",
		"origin":       "https://17.live",
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",
	})

	c := &Client{
		conn:     conn,
		client:   restyclient,
		channels: config.Channels,
	}

	// Add context with timeout for login
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := c.authenticate(ctx, config.Username, config.Password)
	if err != nil {
		c.Disconnect() // Clean up WebSocket connection on failure
		return nil, errors.Wrap(err, "failed to login")
	}

	c.UserProfile = response.UserInfo
	c.client.SetHeader("Authorization", "Bearer "+response.JwtAccessToken)

	return c, nil
}

// Start begins listening for messages and triggers the OnMessage handler
func (c *Client) Connect() error {
	for _, channel := range c.channels {
		if err := c.subscribe(channel); err != nil {
			log.Printf("failed to subscribe to channel %d: %v", channel, err)
			continue
		}
		log.Printf("Connected to: %d", channel)
	}

	go func() {
		defer c.Disconnect() // Ensure cleanup on goroutine exit
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					log.Printf("WebSocket closed unexpectedly: %v", err)
				}
				return
			}
			HandlePacket(c, message, c.onMessage)
		}
	}()
	return nil
}

// Close terminates the WebSocket connection
func (c *Client) Disconnect() {
	c.conn.Close()
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnMessage(handler func(client *Client, message Message)) {
	c.onMessage = handler
}

// Subscribe sends a subscription message to a specific channel
func (c *Client) subscribe(channel int) error {
	subscribeMessage := fmt.Sprintf(`{
		"action": 10,
		"channel": "%s"
	}`, strconv.Itoa(channel))

	return c.conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
}

// Client represents a WebSocket client with a message handler
type Client struct {
	conn        *websocket.Conn
	channels    []int
	client      *resty.Client
	UserProfile UserInfo
	onMessage   func(*Client, Message)
}
