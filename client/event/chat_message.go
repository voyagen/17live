package event

// ChatMessage represents a chat message packet
type ChatMessage struct {
	RoomID    string
	UserID    string
	Username  string
	Picture   string
	Text      string // Avoids conflict with Content() method
	Timestamp int64
}

func (m *ChatMessage) Type() string {
	return "chat"
}
