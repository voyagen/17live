package event

import (
	"encoding/json"
	"fmt"
)

// Parser defines the interface for parsing a specific packet type
type Parser interface {
	Parse(response Response, rawData json.RawMessage) (Packet, error)
}

// Response represents the outer WebSocket response structure
type Response struct {
	Action    int       `json:"action"`
	Messages  []Message `json:"messages"`
	Channel   string    `json:"channel"`
	Timestamp int64     `json:"timestamp"`
}

// parserRegistry holds registered parsers by packet type
var parserRegistry = map[int]Parser{
	3:  &ChatMessageParser{},
	51: &RedEnvelopeParser{},
}

// NewPacket parses raw WebSocket data into a Packet
func NewPacket(responseData []byte) (Packet, error) {
	if len(responseData) == 0 {
		return nil, fmt.Errorf("empty packet")
	}

	var response Response
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if response.Action != 15 || len(response.Messages) == 0 {
		return nil, fmt.Errorf("invalid response: action=%d, messages=%d", response.Action, len(response.Messages))
	}

	data, err := decompressGzip(response.Messages[0].Data)
	if err != nil {
		return nil, fmt.Errorf("decompress data: %w", err)
	}

	packetType, err := packetType(data)
	if err != nil {
		return nil, fmt.Errorf("retrieve packet type: %w", err)
	}

	// check if packet type is supported
	parser, ok := parserRegistry[packetType]
	if !ok {
		return nil, fmt.Errorf("unsupported packet type: %d", packetType)
	}

	var rawData json.RawMessage
	if err := json.Unmarshal([]byte(data), &rawData); err != nil {
		return nil, fmt.Errorf("unmarshal raw data: %w", err)
	}

	return parser.Parse(response, rawData)
}

// ChatMessageParser parses chat message packets
type ChatMessageParser struct{}

func (p *ChatMessageParser) Parse(response Response, rawData json.RawMessage) (Packet, error) {
	var data struct {
		CommentMsg struct {
			DisplayUser struct {
				UserID      string `json:"userID"`
				DisplayName string `json:"displayName"`
				Picture     string `json:"picture"`
			} `json:"displayUser"`
			Content string `json:"content"`
		} `json:"commentMsg"`
	}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return nil, fmt.Errorf("unmarshal chat message: %w", err)
	}

	return &ChatMessage{
		RoomID:    response.Channel,
		UserID:    data.CommentMsg.DisplayUser.UserID,
		Username:  data.CommentMsg.DisplayUser.DisplayName,
		Picture:   data.CommentMsg.DisplayUser.Picture,
		Text:      data.CommentMsg.Content,
		Timestamp: response.Timestamp,
	}, nil
}

// RedEnvelopeParser parses red envelope packets
type RedEnvelopeParser struct{}

func (p *RedEnvelopeParser) Parse(response Response, rawData json.RawMessage) (Packet, error) {
	var data struct {
		RedEnvelopeEventInfo struct {
			ID        string `json:"redEnvelopeID"`
			Count     int    `json:"count"`
			StartTime int    `json:"startTime"`
			EndTime   int    `json:"endTime"`
		} `json:"redEnvelopeEventInfo"`
	}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return nil, fmt.Errorf("unmarshal red envelope: %w", err)
	}

	return &RedEnvelopeInfo{
		ID:        data.RedEnvelopeEventInfo.ID,
		RoomID:    response.Channel,
		Count:     data.RedEnvelopeEventInfo.Count,
		StartTime: data.RedEnvelopeEventInfo.StartTime,
		EndTime:   data.RedEnvelopeEventInfo.EndTime,
	}, nil
}
