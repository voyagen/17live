package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// HandleOnMessage processes incoming WebSocket messages and triggers the provided handler.
// It expects a JSON-encoded message, decompresses it if necessary, and invokes the handler with the parsed Message.
func HandlePacket(client *Client, packet []byte, handler func(*Client, Message)) error {
	if len(packet) == 0 {
		return fmt.Errorf("Empty message received")
	}

	if handler == nil {
		return nil
	}

	var response Response
	if err := json.Unmarshal(packet, &response); err != nil {
		return fmt.Errorf("unmarshal message: %w", err)
	}

	if response.Action != 15 || len(response.Messages) == 0 {
		return fmt.Errorf("No messages in response: action: %s", strconv.Itoa(response.Action))
	}

	data, err := decompressGzip(response.Messages[0].Data)
	if err != nil {
		return fmt.Errorf("Failed to decompress message data: %w", err)
	}

	responsetype, err := retrievePacketType(data)
	if err != nil {
		return fmt.Errorf("Failed to retrieve packet type: %w", err)
	}

	var rawData MessageRawData
	if err := json.Unmarshal([]byte(data), &rawData); err != nil {
		return fmt.Errorf("unmarshal message data: %w", err)
	}

	switch responsetype {
	case 3:
		handler(client, Message{
			Channel:   response.Channel,
			UserID:    rawData.CommentMsg.DisplayUser.UserID,
			Username:  rawData.CommentMsg.DisplayUser.DisplayName,
			Picture:   rawData.CommentMsg.DisplayUser.Picture,
			Content:   rawData.CommentMsg.Content,
			Timestamp: response.Timestamp,
		})
		return nil
		// TODO: implement Red Envelope
		// case 51:
		// 	handler(client, RedEnvelope{
		// 		StartTime: rawData.RedEnvelopeEventInfo.StartTime,
		// 		EndTime:   rawData.RedEnvelopeEventInfo.EndTime,
		// 	})
	}
	return nil
}

func HandleOnRedEnvelope(client *Client, data []byte, handler func(*Client, Message)) error {
	return errors.New("Not implemented")
}

func HandleOnSnackReady(client *Client, data []byte, handler func(*Client, Message)) error {
	return errors.New("Not implemented")
}

func HandleOnUserJoined(client *Client, data []byte, handler func(*Client, Message)) error {
	return errors.New("Not implemented")
}
