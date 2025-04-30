package client

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

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

	payload := struct {
		OpenID   string `json:"openID"`
		Password string `json:"password"`
		Language string `json:"language"`
	}{
		OpenID:   username,
		Password: hashedPassword,
		Language: "EN",
	}

	// Send request with context and timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
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

	// Check if Data is a JSON-encoded string
	var dataString string
	if err := json.Unmarshal(loginResponse.Data, &dataString); err == nil && dataString != "" {
		var decodedData json.RawMessage
		if err := json.Unmarshal([]byte(dataString), &decodedData); err != nil {
			return nil, errors.Wrapf(err, "failed to decode JSON string in Data: %s", dataString)
		}
		loginResponse.Data = decodedData // Update Data with decoded JSON
	}

	// Decode successful response
	var dataResponse LoginResponseData
	if err := json.Unmarshal(loginResponse.Data, &dataResponse); err != nil {
		return nil, errors.Wrapf(err, "failed to decode login response data: %s", string(loginResponse.Data))
	}

	// Validate critical fields
	if dataResponse.JwtAccessToken == "" || dataResponse.UserInfo.UserID == "" {
		return nil, errors.New("invalid login response: missing token or user ID")
	}

	return &dataResponse, nil
}
