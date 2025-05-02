package websocket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/voyagen/17live/client/event"
)

const (
	WebsocketURL = "wss://17media-realtime.ably.io/?key=qvDtFQ.0xBeRA:iYWpd3nD2QHE6Sjm&format=json&heartbeats=true&v=1.1&lib=js-web-1.1.25"
)

// Websocket manages the WebSocket connection
type Websocket struct {
	conn *websocket.Conn
}

// NewWebsocket creates a new WebSocket connection
func NewWebsocket() (*Websocket, error) {
	conn, _, err := websocket.DefaultDialer.Dial(WebsocketURL, http.Header{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to websocket: %v", err)
	}
	return &Websocket{conn: conn}, nil
}

func (w *Websocket) Join(channelID int) error {
	payload := fmt.Sprintf(`{
		"action": 10,
		"channel": "%d"
	}`, channelID)
	return w.conn.WriteMessage(websocket.TextMessage, []byte(payload))
}

// ReadPackets reads WebSocket messages and sends them to a channel
func (w *Websocket) ReadPackets(ctx context.Context, messages chan<- []byte) error {
	defer w.conn.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, data, err := w.conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("failed to read websocket message: %v", err)
			}
			messages <- data
		}
	}
}

// Close closes the WebSocket connection
func (w *Websocket) Close() error {
	return w.conn.Close()
}

// PacketProcessor processes raw messages into packets
func PacketProcessor(ctx context.Context, messages <-chan []byte, packets chan<- event.Packet) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-messages:
			packet, err := event.NewPacket(data)
			if err != nil {
				// Log error and continue
				continue
			}
			packets <- packet
		}
	}
}
