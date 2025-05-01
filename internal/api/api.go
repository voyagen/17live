package api

import "github.com/go-resty/resty/v2"

// TODO
type Client struct {
	Client *resty.Client
}

func NewClient(token string) (*Client, error) {
	client := &Client{}
	return client, nil
}
