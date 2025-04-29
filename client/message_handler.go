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
func HandleOnMessage(client *Client, data []byte, handler func(*Client, Message)) error {
	if len(data) == 0 {
		return fmt.Errorf("Empty message received")
	}

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return fmt.Errorf("unmarshal message: %w", err)
	}

	if response.Action != ActionMessage {
		return nil // Silently ignore unsupported actions
	}

	if len(response.Messages) == 0 {
		slog.Warn("No messages in response", "action", response.Action)
		return nil
	}

	// Process the first message
	decompressedPayload, err := decompressGzip(response.Messages[0].Data)
	if err != nil {
		slog.Error("Failed to decompress message data", "error", err)
		return fmt.Errorf("decompress data: %w", err)
	}

	var rawData MessageRawData
	if err := json.Unmarshal([]byte(decompressedPayload), &rawData); err != nil {
		slog.Error("Failed to unmarshal message data", "error", err, "raw", string(decompressedPayload))
		return fmt.Errorf("unmarshal message data: %w", err)
	}

	// TODO: Accept more payloads
	if rawData.Type != 3 {
		// fmt.Println(decompressedPayload)
		return nil
	}

	if handler == nil {
		slog.Warn("No handler provided for message", "channel", rawData.CommentMsg.StreamerUserID)
		return nil
	}

	// Construct and invoke handler with the message
	payload := Message{
		Channel:  response.Channel,
		UserID:   rawData.CommentMsg.DisplayUser.UserID,
		Username: rawData.CommentMsg.DisplayUser.DisplayName,
		Picture:  rawData.CommentMsg.DisplayUser.Picture,
		Content:  rawData.CommentMsg.Content,
		Tags:     []string{},
	}

	handler(client, payload)
	return nil
}
