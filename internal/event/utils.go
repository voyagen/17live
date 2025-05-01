package event

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

// decompressGzip decompresses a Base64-encoded GZIP string
func decompressGzip(encodedStr string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	reader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return "", fmt.Errorf("gzip reader: %w", err)
	}
	defer reader.Close()

	var decompressedData bytes.Buffer
	if _, err := io.Copy(&decompressedData, reader); err != nil {
		return "", fmt.Errorf("decompress: %w", err)
	}

	return decompressedData.String(), nil
}

// packetType extracts the packet type from raw JSON data
func packetType(data string) (int, error) {
	var msg struct {
		Type int `json:"type"`
	}
	if err := json.Unmarshal([]byte(data), &msg); err != nil {
		return -1, fmt.Errorf("unmarshal packet type: %w", err)
	}
	return msg.Type, nil
}
