package client

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Login authenticates a user with the provided credentials
func (c *Client) authenticate(ctx context.Context, username, password string) (*LoginResponseData, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	// Note: MD5 is weak; consider stronger hashing if API supports it
	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	// Use struct for type safety and consistency with headers
	payload := struct {
		OpenID   string `json:"openID"`
		Password string `json:"password"`
		Language string `json:"language"`
	}{
		OpenID:   username,
		Password: hashedPassword,
		Language: "EN",
	}

	// Send request with context
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(payload).
		Post("https://wap-api.17app.co/api/v1/auth/loginAction")
	if err != nil {
		return nil, errors.Wrap(err, "failed to send login request")
	}

	// Handle non-200 status codes
	if resp.StatusCode() != http.StatusOK {
		body, _ := io.ReadAll(resp.RawResponse.Body)
		return nil, errors.Errorf("login failed: status code %d, body: %s", resp.StatusCode(), string(body))
	}

	// Decode response
	var loginResponse LoginResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&loginResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response")
	}

	// Check for error response
	var errorResponse LoginErrorResponse
	if err := json.Unmarshal(loginResponse.Data, &errorResponse); err == nil && errorResponse.Result == "fail" {
		return nil, errors.New(errorResponse.Message)
	}

	// Decode successful response
	var dataResponse LoginResponseData
	if err := json.Unmarshal(loginResponse.Data, &dataResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode login response data")
	}

	// Validate critical fields
	if dataResponse.JwtAccessToken == "" || dataResponse.UserInfo.UserID == "" {
		return nil, errors.New("invalid login response: missing token or user ID")
	}

	return &dataResponse, nil
}

// LoginResponse represents the API response structure
type LoginResponse struct {
	Data json.RawMessage `json:"data"`
}

// LoginErrorResponse holds error response data
type LoginErrorResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type LoginResponseData struct {
	UserInfo             UserInfo `json:"userInfo"`
	Message              string   `json:"message"`
	Result               string   `json:"result"`
	RefreshToken         string   `json:"refreshToken"`
	JwtAccessToken       string   `json:"jwtAccessToken"`
	AccessToken          string   `json:"accessToken"`
	GiftModuleState      int      `json:"giftModuleState"`
	Word                 string   `json:"word"`
	AbtestNewbieFocus    string   `json:"abtestNewbieFocus"`
	AbtestNewbieGuidance string   `json:"abtestNewbieGuidance"`
	AbtestNewbieGuide    string   `json:"abtestNewbieGuide"`
	ShowRecommend        bool     `json:"showRecommend"`
	AutoEnterLive        struct {
		Auto         bool `json:"auto"`
		LiveStreamID int  `json:"liveStreamID"`
	} `json:"autoEnterLive"`
	NewbieEnhanceGuidanceStyle       int  `json:"newbieEnhanceGuidanceStyle"`
	NewbieGuidanceFocusMissionEnable bool `json:"newbieGuidanceFocusMissionEnable"`
}
