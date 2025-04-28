package client

type LoginResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"`
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

// --- LiveStreamerData struct ---
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
