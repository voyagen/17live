package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// SeventeenLiveClient is a client for interacting with the 17Live API.
type SeventeenLiveClient struct {
	client      *resty.Client
	UserProfile UserInfo
}

// NewSeventeenLiveClient creates a new SeventeenLiveClient with the provided login details.
func NewClient(username string, password string) (*SeventeenLiveClient, error) {
	if username == "" {
		return nil, errors.New("Username cannot be empty")
	}

	if password == "" {
		return nil, errors.New("Password must be positive")
	}

	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"devicetype":   "WEB",
		"language":     "GLOBAL",
		"origin":       "https://17.live",
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",
	})

	c := &SeventeenLiveClient{
		client:      client,
		UserProfile: UserInfo{},
	}

	response, err := c.login(username, password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to login")
	}
	c.UserProfile = response.UserInfo
	c.client.SetHeader("Authorization", "Bearer "+response.JwtAccessToken)

	return c, nil
}

func (c *SeventeenLiveClient) login(username string, password string) (*LoginResponseData, error) {
	// Generate MD5 hash
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])

	payload := map[string]string{
		"openID":   username,
		"password": password,
		"language": "EN",
	}

	resp, err := c.client.R().SetBody(payload).Post("https://wap-api.17app.co/api/v1/auth/loginAction")
	if err != nil {
		return nil, errors.Wrap(err, "failed to send login request")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.Errorf("login failed: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	loginResponse := LoginResponse{}
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&loginResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response")
	}

	var dataResponse LoginResponseData
	if err := json.Unmarshal([]byte(loginResponse.Data), &dataResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response data")
	}

	return &dataResponse, nil
}
