package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/voyagen/17live/internal/event"
)

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnMessage(handler func(client *Client, message *event.ChatMessage)) {
	c.onMessage = handler
}

// SetOnMessage sets the callback function for handling incoming messages
func (c *Client) OnRedEnvelopeInfo(handler func(client *Client, message *event.RedEnvelopeInfo)) {
	c.onRedEnvelopeInfo = handler
}

// SendMessage sends a message
func (c *Client) SendMessage(roomID int, comment string) (*resty.Response, error) {
	return c.api.SendMessage(roomID, comment)
}

// PokeAll sends a poke request to all the users in the livestream.
func (c *Client) PokeAll(roomID int) (*resty.Response, error) {
	return c.api.PokeAll(roomID)
}

// Poke sends a poke request to the specified user.
func (c *Client) Poke(userID string, roomID int) (*resty.Response, error) {
	return c.api.Poke(userID, roomID)
}

// Poke sends a poke request to the specified user back.
func (c *Client) PokeBack(roomID int) (*resty.Response, error) {
	return c.api.PokeBack(roomID)
}

// ShareFacebook sends a Facebook share reaction.
func (c *Client) ShareFacebook(roomID int) (*resty.Response, error) {
	return c.api.ShareFacebook(roomID)
}

// Share17Live sends a 17Live share reaction.
func (c *Client) Share17Live(roomID int) (*resty.Response, error) {
	return c.api.Share17Live(roomID)
}

// Like sends a like reaction.
func (c *Client) Like(roomID int) (*resty.Response, error) {
	return c.api.Like(roomID)
}

// Follow sends a follow request to the streamer
func (c *Client) Follow(roomID int) (*resty.Response, error) {
	return c.api.Follow(roomID)
}

// Unfollow sends a unfollow request to the streamer
func (c *Client) Unfollow(roomID int) (*resty.Response, error) {
	return c.api.Unfollow(roomID)
}
