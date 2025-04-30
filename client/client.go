package client

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	response, err := c.login(ctx, config.Username, config.Password)
	if err != nil {
		c.Disconnect() // Clean up WebSocket connection on failure
		return nil, errors.Wrap(err, "failed to login")
	}

	c.UserProfile = response.UserInfo
	c.client.SetHeader("Authorization", "Bearer "+response.JwtAccessToken)

	return c, nil
}

// Login authenticates a user with the provided credentials
func (c *Client) login(ctx context.Context, username, password string) (*LoginResponseData, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	// Note: MD5 is weak; consider stronger hashing if API supports it
	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	// Use struct for type safety and consistency with headers
	payload := struct {
		OpenID   string `json:"openID"`
		Password string `json:"password"`
		Language string `json:"language"`
	}{
		OpenID:   username,
		Password: hashedPassword,
		Language: "EN",
	}

	// Send request with context
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(payload).
		Post("https://wap-api.17app.co/api/v1/auth/loginAction")
	if err != nil {
		return nil, errors.Wrap(err, "failed to send login request")
	}

	// Handle non-200 status codes
	if resp.StatusCode() != http.StatusOK {
		body, _ := io.ReadAll(resp.RawResponse.Body)
		return nil, errors.Errorf("login failed: status code %d, body: %s", resp.StatusCode(), string(body))
	}

	// Decode response
	var loginResponse LoginResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&loginResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response")
	}

	// Check for error response
	var errorResponse LoginErrorResponse
	if err := json.Unmarshal(loginResponse.Data, &errorResponse); err == nil && errorResponse.Result == "fail" {
		return nil, errors.New(errorResponse.Message)
	}

	// Decode successful response
	var dataResponse LoginResponseData
	if err := json.Unmarshal(loginResponse.Data, &dataResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response data")
	}

	// Validate critical fields
	if dataResponse.JwtAccessToken == "" || dataResponse.UserInfo.UserID == "" {
		return nil, errors.New("invalid login response: missing token or user ID")
	}

	return &dataResponse, nil
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnMessage(handler func(client *Client, message Message)) {
	c.onMessage = handler
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
			HandleOnMessage(c, message, c.onMessage)
		}
	}()
	return nil
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
