package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// fetchStreamerProfile retrieves streamer information for the given room ID.
func (c *SeventeenLiveClient) fetchStreamerProfile(roomID int) (StreamerProfile, error) {
	url := fmt.Sprintf("https://wap-api.17app.co/api/v1/user/room/%d", roomID)

	resp, err := c.client.R().Get(url)
	if err != nil {
		return StreamerProfile{}, errors.Wrap(err, "failed to send request")
	}

	if resp.StatusCode() != http.StatusOK {
		return StreamerProfile{}, errors.Errorf("request failed: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	var profile StreamerProfile
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&profile); err != nil {
		return StreamerProfile{}, errors.Wrap(err, "failed to decode response")
	}

	return profile, nil
}

// SendMessage sends a message
func (c *SeventeenLiveClient) SendMessage(roomID int, comment string) (*resty.Response, error) {
	payload := CommentRequest{
		Comment:     comment,
		CommentType: 0,
	}

	return c.client.R().SetBody(payload).
		Post("https://wap-api.17app.co/api/v1/lives/" + strconv.Itoa(roomID) + "/comments")
}

// Poke sends a poke request to the specified user.
func (c *SeventeenLiveClient) Poke(userID string, roomID int) (*resty.Response, error) {
	payload := PokeRequest{
		UserID:     userID,
		IsPokeBack: false,
		SrcID:      strconv.Itoa(roomID),
	}

	return c.client.R().SetBody(payload).
		Post("https://wap-api.17app.co/api/v1/pokes")
}

// Poke sends a poke request to the specified user back.
func (c *SeventeenLiveClient) PokeBack(roomID int) (*resty.Response, error) {
	streamer, err := c.fetchStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := PokeRequest{
		UserID:     streamer.UserID,
		IsPokeBack: true,
		SrcID:      strconv.Itoa(streamer.RoomID),
	}

	return c.client.R().SetBody(payload).
		Post("https://wap-api.17app.co/api/v1/pokes")
}

// ShareFacebook sends a Facebook share reaction.
func (c *SeventeenLiveClient) ShareFacebook(roomID int) (*resty.Response, error) {
	streamer, err := c.fetchStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         0,
	}

	return c.client.R().SetBody(payload).
		Post(fmt.Sprintf("https://wap-api.17app.co/api/v1/lives/%d/reacts", streamer.RoomID))
}

// Share17Live sends a 17Live share reaction.
func (c *SeventeenLiveClient) Share17Live(roomID int) (*resty.Response, error) {
	streamer, err := c.fetchStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         1,
	}

	return c.client.R().SetBody(payload).
		Post(fmt.Sprintf("https://wap-api.17app.co/api/v1/lives/%d/reacts", streamer.RoomID))
}

// Follow sends a follow request to the streamer
func (c *SeventeenLiveClient) Follow(roomID int) (*resty.Response, error) {
	streamer, err := c.fetchStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	followPayload := FollowRequest{
		UserID: streamer.UserID,
	}

	_, err = c.client.R().SetBody(followPayload).
		Post(fmt.Sprintf("https://wap-api.17app.co/api/v1/follow/users/%s", streamer.UserID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to follow user")
	}

	reactionPayload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         3,
	}

	return c.client.R().SetBody(reactionPayload).
		Post(fmt.Sprintf("https://wap-api.17app.co/api/v1/lives/%d/reacts", streamer.RoomID))
}

// Unfollow sends a unfollow request to the streamer
func (c *SeventeenLiveClient) Unfollow(roomID int) (*resty.Response, error) {
	streamer, err := c.fetchStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	followPayload := FollowRequest{
		UserID: streamer.UserID,
	}

	return c.client.R().SetBody(followPayload).
		Delete(fmt.Sprintf("https://wap-api.17app.co/api/v1/follow/users/%s", streamer.UserID))
}
