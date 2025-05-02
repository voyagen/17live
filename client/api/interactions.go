package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// Constants for URLs
const (
	BaseURL          = "https://wap-api.17app.co/api/v1"
	EnterEndpoint    = "lives/%d/enter?"
	CommentsEndpoint = "/lives/%d/comments"
	PokeAllEndpoint  = "/pokes/pokeAll"
	PokeEndpoint     = "/pokes"
	ReactsEndpoint   = "/lives/%d/reacts"
	FollowEndpoint   = "/follow/users/%s"
	RoomEndpoint     = "/user/room/%d"
)

// Constants for static values
const (
	CommentTypeDefault        = 0
	ReceiverGroupAll          = 2
	ReactionTypeFacebookShare = 0
	ReactionType17LiveShare   = 1
	ReactionTypeLike          = 2
	ReactionTypeFollow        = 3
)

// Enter a room
func (c *Client) Enter(roomID int) (*resty.Response, error) {
	return c.Client.R().
		Post(fmt.Sprintf(BaseURL+EnterEndpoint, roomID))
}

// SendMessage sends a message
func (c *Client) SendMessage(roomID int, comment string) (*resty.Response, error) {
	payload := CommentRequest{
		Comment:     comment,
		CommentType: CommentTypeDefault,
	}

	return c.Client.R().SetBody(payload).
		Post(fmt.Sprintf(BaseURL+CommentsEndpoint, roomID))
}

// PokeAll sends a poke request to all the users in the livestream.
func (c *Client) PokeAll(roomID int) (*resty.Response, error) {
	payload := PokeAllRequest{
		ReceiverGroup: ReceiverGroupAll,
		LiveStreamID:  strconv.Itoa(roomID),
	}

	return c.Client.R().SetBody(payload).
		Post(BaseURL + PokeAllEndpoint)
}

// Poke sends a poke request to the specified user.
func (c *Client) Poke(userID string, roomID int) (*resty.Response, error) {
	payload := PokeRequest{
		UserID:     userID,
		IsPokeBack: false,
		SrcID:      strconv.Itoa(roomID),
	}

	return c.Client.R().SetBody(payload).
		Post(BaseURL + PokeEndpoint)
}

// PokeBack sends a poke request to the specified user back.
func (c *Client) PokeBack(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := PokeRequest{
		UserID:     streamer.UserID,
		IsPokeBack: true,
		SrcID:      strconv.Itoa(streamer.RoomID),
	}

	return c.Client.R().SetBody(payload).
		Post(BaseURL + PokeEndpoint)
}

// ShareFacebook sends a Facebook share reaction.
func (c *Client) ShareFacebook(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         ReactionTypeFacebookShare,
	}

	return c.Client.R().SetBody(payload).
		Post(fmt.Sprintf(BaseURL+ReactsEndpoint, streamer.RoomID))
}

// Share17Live sends a 17Live share reaction.
func (c *Client) Share17Live(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         ReactionType17LiveShare,
	}

	return c.Client.R().SetBody(payload).
		Post(fmt.Sprintf(BaseURL+ReactsEndpoint, streamer.RoomID))
}

// Like sends a like reaction.
func (c *Client) Like(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         ReactionTypeLike,
	}

	return c.Client.R().SetBody(payload).
		Post(fmt.Sprintf(BaseURL+ReactsEndpoint, streamer.RoomID))
}

// Follow sends a follow request to the streamer
func (c *Client) Follow(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	followPayload := FollowRequest{
		UserID: streamer.UserID,
	}

	_, err = c.Client.R().SetBody(followPayload).
		Post(fmt.Sprintf(BaseURL+FollowEndpoint, streamer.UserID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to follow user")
	}

	payload := ReactionRequest{
		UserID:       streamer.UserID,
		LiveStreamID: streamer.RoomID,
		Type:         ReactionTypeFollow,
	}

	return c.Client.R().SetBody(payload).
		Post(fmt.Sprintf(BaseURL+ReactsEndpoint, streamer.RoomID))
}

// Unfollow sends a unfollow request to the streamer
func (c *Client) Unfollow(roomID int) (*resty.Response, error) {
	streamer, err := c.getStreamerProfile(roomID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch streamer profile")
	}

	followPayload := FollowRequest{
		UserID: streamer.UserID,
	}

	return c.Client.R().SetBody(followPayload).
		Delete(fmt.Sprintf(BaseURL+FollowEndpoint, streamer.UserID))
}

// fetchStreamerProfile retrieves streamer information for the given room ID.
func (c *Client) getStreamerProfile(roomID int) (StreamerProfile, error) {
	url := fmt.Sprintf(BaseURL+RoomEndpoint, roomID)

	resp, err := c.Client.R().Get(url)
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

// PokeRequest represents a poke API request payload.
type PokeAllRequest struct {
	ReceiverGroup int    `json:"receiverGroup"`
	LiveStreamID  string `json:"liveStreamID"`
}

// PokeRequest represents a poke API request payload.
type PokeRequest struct {
	UserID     string `json:"userID"`
	IsPokeBack bool   `json:"isPokeBack"`
	SrcID      string `json:"srcID"`
}

// CommentRequest represents a comment API request payload.
type CommentRequest struct {
	Comment     string `json:"comment"`
	CommentType int    `json:"commentType"`
}

// ReactionRequest represents a reaction API request payload.
type ReactionRequest struct {
	UserID       string `json:"userID"`
	LiveStreamID int    `json:"liveStreamID"`
	Type         int    `json:"type"`
}

// FollowRequest represents a follow API request payload.
type FollowRequest struct {
	UserID string `json:"userID"`
}

type StreamerProfile struct {
	UserID                string `json:"userID"`
	OpenID                string `json:"openID"`
	DisplayName           string `json:"displayName"`
	Name                  string `json:"name"`
	Bio                   string `json:"bio"`
	Picture               string `json:"picture"`
	Website               string `json:"website"`
	FollowerCount         int    `json:"followerCount"`
	FollowingCount        int    `json:"followingCount"`
	ReceivedLikeCount     int    `json:"receivedLikeCount"`
	LikeCount             int    `json:"likeCount"`
	IsFollowing           int    `json:"isFollowing"`
	IsNotif               int    `json:"isNotif"`
	IsBlocked             int    `json:"isBlocked"`
	FollowTime            int    `json:"followTime"`
	FollowRequestTime     int    `json:"followRequestTime"`
	RoomID                int    `json:"roomID"`
	PrivacyMode           string `json:"privacyMode"`
	BallerLevel           int    `json:"ballerLevel"`
	PostCount             int    `json:"postCount"`
	IsCelebrity           int    `json:"isCelebrity"`
	Baller                int    `json:"baller"`
	Level                 int    `json:"level"`
	FollowPrivacyMode     int    `json:"followPrivacyMode"`
	RevenueShareIndicator string `json:"revenueShareIndicator"`
	ClanStatus            int    `json:"clanStatus"`
	ClanInfo              struct {
		DisplayClans []interface{} `json:"displayClans"`
		JoinCount    int           `json:"joinCount"`
	} `json:"clanInfo"`
	BadgeInfo []struct {
		BadgeID   string `json:"badgeID"`
		BadgeName struct {
			Key string `json:"key"`
		} `json:"badgeName"`
		Description struct {
			Key string `json:"key"`
		} `json:"description"`
	} `json:"badgeInfo"`
}
