package client

import (
	"encoding/json"
	"log"
)

// HandleMessage processes incoming WebSocket messages and triggers the handler
func HandleMessage(message []byte, handler func(msg DecryptedMessage)) {
	var response Response
	err := json.Unmarshal(message, &response)
	if err != nil {
		log.Printf("Error unmarshalling message: %v, raw: %s\n", err, string(message))
		return
	}

	if response.Action == 15 && len(response.Messages) > 0 {
		decompressed, err := decompressGzip(response.Messages[0].Data)
		if err != nil {
			log.Printf("Error decompressing data: %v\n", err)
			return
		}

		var decrypted DecryptedMessage
		err = json.Unmarshal([]byte(decompressed), &decrypted)
		if err != nil {
			log.Printf("Error unmarshalling message: %v, raw: %s\n", err, string(decompressed))
			return
		}

		// Trigger the handler if set
		if handler != nil {
			handler(decrypted)
		}
	}
}
