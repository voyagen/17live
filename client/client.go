package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// NewClient creates a new WebSocket client
func NewClient(config ClientConfig) (*Client, error) {
	if config.Username == "" {
		return nil, errors.New("Config username cannot be empty")
	}

	if config.Password == "" {
		return nil, errors.New("Config password cannot be empty")
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

	client := &Client{
		conn:     conn,
		client:   restyclient,
		channels: config.Channels,
	}

	response, err := client.login(config.Username, config.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to login")
	}
	client.UserProfile = response.UserInfo
	client.client.SetHeader("Authorization", "Bearer "+response.JwtAccessToken)

	return client, nil
}

func (c *Client) login(username string, password string) (*LoginResponseData, error) {
	// Generate MD5 hash
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])

	payload := map[string]string{
		"openID":   username,
		"password": password,
		"language": "EN",
	}

	resp, err := c.client.R().SetBody(payload).Post("https://wap-api.17app.co/api/v1/auth/loginAction")
	if err != nil {
		return nil, errors.Wrap(err, "failed to send login request")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.Errorf("login failed: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	loginResponse := LoginResponse{}
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&loginResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response")
	}

	var dataResponse LoginResponseData
	if err := json.Unmarshal([]byte(loginResponse.Data), &dataResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response data")
	}

	return &dataResponse, nil
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnMessage(handler func(client *Client, message Message)) {
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
			HandleOnMessage(c, message, c.onMessage)
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
