package client

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/voyagen/17live/internal/api"
	"github.com/voyagen/17live/internal/auth"
	"github.com/voyagen/17live/internal/event"
	"github.com/voyagen/17live/internal/websocket"
)

// Client represents a WebSocket client with a message handler
type Client struct {
	conn *websocket.Websocket
	api  *api.Client

	User auth.UserInfo

	onMessage         func(*Client, *event.ChatMessage)
	onRedEnvelopeInfo func(*Client, *event.RedEnvelopeInfo)
}

type Config struct {
	Username string
	Password string
	Channels []int
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

	return &Client{
		conn: conn,
		api:  apiClient,
		User: user.UserInfo,
	}, nil
}

// Connect starts reading packets and dispatches them to the appropriate handlers
func (c *Client) Connect() error {
	if c.conn == nil {
		return errors.New("websocket connection is nil")
	}

	handler := func(p event.Packet) {
		switch pkt := p.(type) {
		case *event.ChatMessage:
			if c.onMessage != nil {
				c.onMessage(c, pkt)
			}
		case *event.RedEnvelopeInfo:
			if c.onRedEnvelopeInfo != nil {
				c.onRedEnvelopeInfo(c, pkt)
			}
		default:
			log.Printf("Unknown packet type: %+v", pkt)
		}
	}

	if err := c.conn.ReadPackets(handler); err != nil {
		return errors.Wrap(err, "error reading packets")
	}

	return nil
}
