package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// TODO
type Client struct {
	Client *resty.Client
}

func NewClient(token string) (*Client, error) {
	r := resty.New().
		SetRetryCount(3).
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token),
			"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",
			"Origin":        "https://17.live",
			"Referer":       "https://17.live/",
			"devicetype":    "WEB",
			"language":      "GLOBAL",
		})

	return &Client{
		Client: r,
	}, nil
}
