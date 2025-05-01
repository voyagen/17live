package websocket

import (
	"fmt"
	"net/http"

	"github.com/voyagen/17live/internal/event"

	"github.com/gorilla/websocket"
)

const (
	WebsocketURL = "wss://17media-realtime.ably.io/?key=qvDtFQ.0xBeRA:iYWpd3nD2QHE6Sjm&format=json&heartbeats=true&v=1.1&lib=js-web-1.1.25"
)

// Websocket manages the WebSocket connection
type Websocket struct {
	conn *websocket.Conn
}

// NewWebsocket creates a new WebSocket connection with JWT authentication
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

// ReadPackets reads and processes WebSocket packets
func (w *Websocket) ReadPackets(handler func(event.Packet)) error {
	for {
		_, data, err := w.conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("failed to read websocket message: %v", err)
		}
		packet, err := event.NewPacket(data)
		if err != nil {
			// Failed to parse packet
			continue
		}
		handler(packet)
	}
}

// Close closes the WebSocket connection
func (w *Websocket) Close() error {
	return w.conn.Close()
}
