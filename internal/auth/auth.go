package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	loginURL      = "https://wap-api.17app.co/api/v1/auth/loginAction"
	defaultLang   = "EN"
	deviceType    = "WEB"
	contentType   = "application/json"
	defaultOrigin = "https://17.live"
	userAgent     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36"
)

// Auth manages authentication with the 17live API
type Auth struct {
	client *resty.Client
}

// New creates a new Auth instance
func New() *Auth {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": contentType,
		"devicetype":   deviceType,
		"language":     "GLOBAL",
		"origin":       defaultOrigin,
		"user-agent":   userAgent,
	})
	return &Auth{client: client}
}

// UserInfo holds user profile information
type UserInfo struct {
	UserID                        string        `json:"userID"`
	OpenID                        string        `json:"openID"`
	DisplayName                   string        `json:"displayName"`
	Name                          string        `json:"name"`
	Bio                           string        `json:"bio"`
	Picture                       string        `json:"picture"`
	Website                       string        `json:"website"`
	FollowerCount                 int           `json:"followerCount"`
	FollowingCount                int           `json:"followingCount"`
	ReceivedLikeCount             int           `json:"receivedLikeCount"`
	LikeCount                     int           `json:"likeCount"`
	IsFollowing                   int           `json:"isFollowing"`
	IsNotif                       int           `json:"isNotif"`
	IsBlocked                     int           `json:"isBlocked"`
	FollowTime                    int           `json:"followTime"`
	FollowRequestTime             int           `json:"followRequestTime"`
	RoomID                        int           `json:"roomID"`
	PrivacyMode                   string        `json:"privacyMode"`
	BallerLevel                   int           `json:"ballerLevel"`
	PostCount                     int           `json:"postCount"`
	IsCelebrity                   int           `json:"isCelebrity"`
	Baller                        int           `json:"baller"`
	Level                         int           `json:"level"`
	FollowPrivacyMode             int           `json:"followPrivacyMode"`
	RevenueShareIndicator         string        `json:"revenueShareIndicator"`
	ClanStatus                    int           `json:"clanStatus"`
	BadgeInfo                     []interface{} `json:"badgeInfo"`
	Region                        string        `json:"region"`
	HideAllPointToLeaderboard     int           `json:"hideAllPointToLeaderboard"`
	EnableShop                    int           `json:"enableShop"`
	MonthlyVIPBadges              struct{}      `json:"monthlyVIPBadges"`
	LastLiveTimestamp             int           `json:"lastLiveTimestamp"`
	LastCreateLiveTimestamp       int           `json:"lastCreateLiveTimestamp"`
	LastLiveRegion                string        `json:"lastLiveRegion"`
	LoyaltyInfo                   []interface{} `json:"loyaltyInfo"`
	StreamerRecapEnable           bool          `json:"streamerRecapEnable"`
	GloryroadMode                 int           `json:"gloryroadMode"`
	LastUsedHashtags              []interface{} `json:"lastUsedHashtags"`
	NewbieDisplayAllGiftTabsToast bool          `json:"newbieDisplayAllGiftTabsToast"`
	IsUnderaged                   bool          `json:"isUnderaged"`
	IsEmailVerified               int           `json:"isEmailVerified"`
	CommentShadowColor            string        `json:"commentShadowColor"`
	IsFreePrivateMsgEnabled       bool          `json:"isFreePrivateMsgEnabled"`
	IsVliverOnlyModeEnabled       bool          `json:"isVliverOnlyModeEnabled"`
}

// loginPayload is the request body sent to the API
type loginPayload struct {
	OpenID   string `json:"openID"`
	Password string `json:"password"`
	Language string `json:"language"`
}

// LoginResponseData holds successful login data
type LoginResponseData struct {
	JwtAccessToken string   `json:"jwtAccessToken"`
	UserInfo       UserInfo `json:"userInfo"`
}

// apiResponse wraps the API response
type apiResponse struct {
	Data json.RawMessage `json:"data"`
}

// apiError represents an error response format
type apiError struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

// Login authenticates with username and password, returns user info and JWT token
func (a *Auth) Login(ctx context.Context, username, password string) (LoginResponseData, error) {
	if username == "" || password == "" {
		return LoginResponseData{}, errors.New("username and password are required")
	}

	// Hash password with MD5
	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	payload := loginPayload{
		OpenID:   username,
		Password: hashedPassword,
		Language: defaultLang,
	}

	resp, err := a.client.R().
		SetContext(ctx).
		SetBody(payload).
		Post(loginURL)
	if err != nil {
		return LoginResponseData{}, errors.Wrap(err, "failed to send login request")
	}

	if resp.StatusCode() != http.StatusOK {
		return LoginResponseData{}, fmt.Errorf("login failed: status code %d, body: %s", resp.StatusCode(), resp.String())
	}

	var rawResp apiResponse
	if err := json.Unmarshal(resp.Body(), &rawResp); err != nil {
		return LoginResponseData{}, errors.Wrap(err, "failed to decode API response")
	}

	// First check if it contains an error structure
	var errResp apiError
	if err := json.Unmarshal(rawResp.Data, &errResp); err == nil && errResp.Result == "fail" {
		return LoginResponseData{}, errors.New(errResp.Message)
	}

	// Decode successful login data
	// Decode successful login data
	var loginData LoginResponseData

	// Step 1: Unmarshal rawResp.Data into a string (it's a JSON-encoded string)
	var innerJSON string
	if err := json.Unmarshal(rawResp.Data, &innerJSON); err != nil {
		return LoginResponseData{}, errors.Wrap(err, "expected JSON string in 'data' field")
	}

	// Step 2: Unmarshal the inner string into the LoginResponseData struct
	if err := json.Unmarshal([]byte(innerJSON), &loginData); err != nil {
		return LoginResponseData{}, errors.Wrap(err, "failed to decode nested login response data")
	}

	if loginData.JwtAccessToken == "" || loginData.UserInfo.UserID == "" {
		return LoginResponseData{}, errors.New("invalid login response: missing token or user ID")
	}

	return loginData, nil
}
