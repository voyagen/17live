package client

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/voyagen/17live/client/api"
	"github.com/voyagen/17live/client/auth"
	"github.com/voyagen/17live/client/event"
	"github.com/voyagen/17live/client/websocket"
)

// Client represents a WebSocket client with a message handler
type Client struct {
	conn    *websocket.Websocket
	api     *api.Client
	workers int // Store only the number of workers, not the full Config

	User auth.UserInfo

	onMessage         func(*Client, *event.ChatMessage)
	onRedEnvelopeInfo func(*Client, *event.RedEnvelopeInfo)
	onPoke            func(*Client, *event.Poke)
	onUserJoined      func(*Client, *event.UserJoined)
}

type Config struct {
	Username string
	Password string
	Channels []int
	Workers  int // Number of packet processing workers
}

// NewClient creates a new 17Live client
func NewClient(config Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	auth := auth.New()
	user, err := auth.Login(ctx, config.Username, config.Password)
	if err != nil {
		return nil, err
	}

	conn, err := websocket.NewWebsocket()
	if err != nil {
		return nil, errors.New("Failed to connect to Websocket")
	}

	// Join the websocket channels
	for _, channel := range config.Channels {
		conn.Join(channel)
	}

	apiClient, err := api.NewClient(user.JwtAccessToken)
	if err != nil {
		return nil, err
	}

	for _, channel := range config.Channels {
		_, err := apiClient.Enter(channel)
		if err != nil {
			return nil, err
		}
	}

	workers := config.Workers
	if workers <= 0 {
		workers = 4 // Default worker count if not specified
	}

	return &Client{
		conn:    conn,
		api:     apiClient,
		workers: workers, // Store only the worker count
		User:    user.UserInfo,
	}, nil
}

// / Connect starts reading packets and dispatches them to handlers
func (c *Client) Connect(ctx context.Context) error {
	if c.conn == nil {
		return errors.New("websocket connection is nil")
	}

	// Channels for fan-in
	messages := make(chan []byte, 100) // Buffered to avoid blocking
	packets := make(chan event.Packet, 100)

	// Number of worker goroutines for packet processing
	workerCount := c.workers

	// Start reading WebSocket messages
	go func() {
		if err := c.conn.ReadPackets(ctx, messages); err != nil {
			log.Printf("Error reading packets: %v", err)
		}
	}()

	// Start packet processing workers
	for i := 0; i < workerCount; i++ {
		go websocket.PacketProcessor(ctx, messages, packets)
	}

	// Dispatch packets to handlers
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case pkt := <-packets:
				switch p := pkt.(type) {
				case *event.ChatMessage:
					if c.onMessage != nil {
						c.onMessage(c, p)
					}
				case *event.RedEnvelopeInfo:
					if c.onRedEnvelopeInfo != nil {
						c.onRedEnvelopeInfo(c, p)
					}
				case *event.Poke:
					if c.onPoke != nil {
						c.onPoke(c, p)
					}
				case *event.UserJoined:
					if c.onUserJoined != nil {
						c.onUserJoined(c, p)
					}
				default:
					log.Printf("Unknown packet type: %+v", p)
				}
			}
		}
	}()

	return nil
}
