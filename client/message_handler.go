package client

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

// Action constants to avoid magic numbers.
const (
	ActionMessage = 15
)

// HandleOnMessage processes incoming WebSocket messages and triggers the provided handler.
// It expects a JSON-encoded message, decompresses it if necessary, and invokes the handler with the parsed Message.
func HandleOnMessage(message []byte, handler func(msg Message)) error {
	if len(message) == 0 {
		return fmt.Errorf("empty message received")
	}

	var resp Response
	if err := json.Unmarshal(message, &resp); err != nil {
		slog.Error("Failed to unmarshal message", "error", err, "raw", string(message))
		return fmt.Errorf("unmarshal message: %w", err)
	}

	if resp.Action != ActionMessage {
		return nil // Silently ignore unsupported actions
	}

	if len(resp.Messages) == 0 {
		slog.Warn("No messages in response", "action", resp.Action)
		return nil
	}

	// Process the first message
	data, err := decompressGzip(resp.Messages[0].Data)
	if err != nil {
		slog.Error("Failed to decompress message data", "error", err)
		return fmt.Errorf("decompress data: %w", err)
	}

	var rawData MessageRawData
	if err := json.Unmarshal([]byte(data), &rawData); err != nil {
		slog.Error("Failed to unmarshal message data", "error", err, "raw", string(data))
		return fmt.Errorf("unmarshal message data: %w", err)
	}

	if rawData.Type != 3 {
		return nil // Ignore messages of incorrect type
	}

	if handler == nil {
		slog.Warn("No handler provided for message", "channel", rawData.CommentMsg.StreamerUserID)
		return nil
	}

	// Construct and invoke handler with the message
	m := Message{
		Channel:  resp.Channel,
		UserID:   rawData.CommentMsg.DisplayUser.UserID,
		Username: rawData.CommentMsg.DisplayUser.DisplayName,
		Picture:  rawData.CommentMsg.DisplayUser.Picture,
		Content:  rawData.CommentMsg.Content,
		Tags:     []string{},
	}
	handler(m)

	return nil
}
