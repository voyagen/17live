package client

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
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
