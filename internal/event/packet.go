package event

// Packet represents a WebSocket packet in the 17live live chat
type Packet interface {
	Type() string
}

type Message struct {
	Data string `json:"data"`
}
