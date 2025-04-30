package client

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

// decompressGzip decompresses a Base64-encoded GZIP string
func decompressGzip(encodedStr string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}

	reader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	decompressedData := new(bytes.Buffer)
	_, err = io.Copy(decompressedData, reader)
	if err != nil {
		return "", err
	}

	return decompressedData.String(), nil
}

func retrievePacketType(data string) (int, error) {
	var msg MessageRawData
	if err := json.Unmarshal([]byte(data), &msg); err != nil {
		slog.Error("Failed to unmarshal message data", "error", err, "raw", string(data))
		return -1, fmt.Errorf("unmarshal message data: %w", err)
	}
	return msg.Type, nil
}
