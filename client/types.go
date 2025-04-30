package client

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

// Response represents the full WebSocket response
type Response struct {
	Action           int    `json:"action"`
	ID               string `json:"id"`
	ConnectionSerial int    `json:"connectionSerial"`
	Channel          string `json:"channel"`
	ChannelSerial    string `json:"channelSerial"`
	Messages         []struct {
		Data string `json:"data"`
	} `json:"messages"`
	Timestamp int64 `json:"timestamp"`
}

type Message struct {
	Channel   string
	UserID    string
	Username  string
	Picture   string
	Content   string
	Timestamp int64
}

type RedEnvelope struct {
	StartTime int
	EndTime   int
}

type MessageRawData struct {
	Type                int         `json:"type"`
	Version             interface{} `json:"version"`
	NewJoinRequest      interface{} `json:"newJoinRequest"`
	WithdrawJoinRequest interface{} `json:"withdrawJoinRequest"`
	AcceptJoinRequest   interface{} `json:"acceptJoinRequest"`
	RejectJoinRequest   interface{} `json:"rejectJoinRequest"`
	JoinRoom            interface{} `json:"joinRoom"`
	LeaveRoom           interface{} `json:"leaveRoom"`
	Kick                interface{} `json:"kick"`
	KickBySkyeye        interface{} `json:"kickBySkyeye"`
	EnableBlab          interface{} `json:"enableBlab"`
	DisableBlab         interface{} `json:"disableBlab"`
	RequestExpire       interface{} `json:"requestExpire"`
	ParticipantExpire   interface{} `json:"participantExpire"`
	StreamerConnectFail interface{} `json:"streamerConnectFail"`
	CommentMsg          struct {
		Marquee struct {
			Type  int `json:"type"`
			Point int `json:"point"`
		} `json:"marquee"`
		Barrage struct {
			Type        int    `json:"type"`
			Count       int    `json:"count"`
			IsInfinite  bool   `json:"isInfinite"`
			Point       int    `json:"point"`
			Name        string `json:"name"`
			AnimationID string `json:"animationID"`
			ExpireTime  int    `json:"expireTime"`
		} `json:"barrage"`
		Stamp          interface{} `json:"stamp"`
		SuffixNameText interface{} `json:"suffixNameText"`
		IsDirty        bool        `json:"isDirty"`
		IsDirtyUser    bool        `json:"isDirtyUser"`
		IsFraud        bool        `json:"isFraud"`
		Region         string      `json:"region"`
		Type           int         `json:"type"`
		SendTime       int64       `json:"sendTime"`
		Level          int         `json:"level"`
		GloryroadMode  int         `json:"gloryroadMode"`
		GloryroadInfo  struct {
			Point        int    `json:"point"`
			Level        int    `json:"level"`
			IconURL      string `json:"iconURL"`
			BadgeIconURL string `json:"badgeIconURL"`
		} `json:"gloryroadInfo"`
		Rank   int `json:"rank"`
		Avatar struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"avatar"`
		PrefixBadge struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"prefixBadge"`
		Name struct {
			Text        string `json:"text"`
			TextColor   string `json:"textColor"`
			TextSize    int    `json:"textSize"`
			ShadowColor string `json:"shadowColor"`
			TagColor    string `json:"tagColor"`
			IsShadow    bool   `json:"isShadow"`
		} `json:"name"`
		MiddleBadge struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"middleBadge"`
		RoleBadge struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"roleBadge"`
		AttendanceBadge struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"attendanceBadge"`
		Comment struct {
			Text        string `json:"text"`
			TextColor   string `json:"textColor"`
			TextSize    int    `json:"textSize"`
			ShadowColor string `json:"shadowColor"`
			TagColor    string `json:"tagColor"`
			IsShadow    bool   `json:"isShadow"`
		} `json:"comment"`
		Icon          interface{} `json:"icon"`
		TopRightBadge struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"topRightBadge"`
		BackgroundColor string `json:"backgroundColor"`
		Border          struct {
			Type                int    `json:"type"`
			URL                 string `json:"URL"`
			ImageWidth          int    `json:"imageWidth"`
			ImageHeight         int    `json:"imageHeight"`
			CommentCornerRadius int    `json:"commentCornerRadius"`
			BorderWidth         int    `json:"borderWidth"`
			BorderHeight        int    `json:"borderHeight"`
			StyleID             string `json:"styleID"`
			DisplayName         string `json:"displayName"`
			EndTime             int    `json:"endTime"`
			Equipped            bool   `json:"equipped"`
			IsNew               bool   `json:"isNew"`
		} `json:"border"`
		CheckinLevel int `json:"checkinLevel"`
		MLevelBadge  struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"mLevelBadge"`
		PrefixNameText     interface{} `json:"prefixNameText"`
		SendToFirebase     interface{} `json:"sendToFirebase"`
		CommentShadowColor string      `json:"commentShadowColor"`
		StreamerUserID     string      `json:"streamerUserID"`
		Content            string      `json:"content"`
		ColorCode          string      `json:"colorCode"`
		BarrageStyle       bool        `json:"barrageStyle"`
		DisplayUser        struct {
			UserID          string `json:"userID"`
			DisplayName     string `json:"displayName"`
			Picture         string `json:"picture"`
			Level           int    `json:"level"`
			IsVIP           bool   `json:"isVIP"`
			IsGuardian      bool   `json:"isGuardian"`
			CheckinLevel    int    `json:"checkinLevel"`
			Producer        int    `json:"producer"`
			Program         int    `json:"program"`
			BadgeURL        string `json:"badgeURL"`
			PfxBadgeURL     string `json:"pfxBadgeURL"`
			PfxBadgeID      string `json:"pfxBadgeID"`
			TopRightIconURL string `json:"topRightIconURL"`
			TopRightIconID  string `json:"topRightIconID"`
			BgColor         string `json:"bgColor"`
			FgColor         string `json:"fgColor"`
			CircleBadgeURL  string `json:"circleBadgeURL"`
			ArmyRank        int    `json:"armyRank"`
			HasProgram      bool   `json:"hasProgram"`
			IsProducer      bool   `json:"isProducer"`
			IsStreamer      bool   `json:"isStreamer"`
			IsDirty         bool   `json:"isDirty"`
			IsDirtyUser     bool   `json:"isDirtyUser"`
			IsFraud         bool   `json:"isFraud"`
			MLevel          int    `json:"mLevel"`
			CheckinBdgURL   string `json:"checkinBdgURL"`
			CheckinCmtURL   string `json:"checkinCmtURL"`
			GloryroadMode   int    `json:"gloryroadMode"`
			GloryroadInfo   struct {
				Point        int    `json:"point"`
				Level        int    `json:"level"`
				IconURL      string `json:"iconURL"`
				BadgeIconURL string `json:"badgeIconURL"`
			} `json:"gloryroadInfo"`
			LevelBadges []struct {
				StyleID int    `json:"styleID"`
				Level   int    `json:"level"`
				IconURL string `json:"iconURL"`
			} `json:"levelBadges"`
		} `json:"displayUser"`
		PrefixBadges []struct {
			URL     string `json:"URL"`
			StyleID string `json:"styleID"`
		} `json:"prefixBadges"`
	} `json:"commentMsg"`
	SubscriberEnterMsg     interface{} `json:"subscriberEnterMsg"`
	ReactMsg               interface{} `json:"reactMsg"`
	DisplayUserMsg         interface{} `json:"displayUserMsg"`
	GuardianInfo           interface{} `json:"guardianInfo"`
	GiftMsg                interface{} `json:"giftMsg"`
	GlobalLuckyBagMsg      interface{} `json:"globalLuckyBagMsg"`
	LiveStreamBroadcastMsg interface{} `json:"liveStreamBroadcastMsg"`
	GlobalInfoUpdateMsg    interface{} `json:"globalInfoUpdateMsg"`
	GameMsg                interface{} `json:"gameMsg"`
	Liveinfo               interface{} `json:"liveinfo"`
	LiveinfoChange         interface{} `json:"liveinfoChange"`
	PokeInfo               interface{} `json:"pokeInfo"`
	PollInfo               interface{} `json:"pollInfo"`
	MediaMessage           interface{} `json:"mediaMessage"`
	GlobalAnnouncementMsg  interface{} `json:"globalAnnouncementMsg"`
	LiveStreamPromotionMsg interface{} `json:"liveStreamPromotionMsg"`
	TriviaMsg              interface{} `json:"triviaMsg"`
	RedEnvelopeEventInfo   struct {
		DisplayInfo struct {
			UserID        string      `json:"userID"`
			DisplayName   string      `json:"displayName"`
			Picture       string      `json:"picture"`
			Name          string      `json:"name"`
			Level         int         `json:"level"`
			OpenID        string      `json:"openID"`
			Region        string      `json:"region"`
			GloryroadInfo interface{} `json:"gloryroadInfo"`
			GloryroadMode int         `json:"gloryroadMode"`
			OnliveInfo    interface{} `json:"onliveInfo"`
			LevelBadges   interface{} `json:"levelBadges"`
		} `json:"displayInfo"`
		StartTime                    int           `json:"startTime"`
		EndTime                      int           `json:"endTime"`
		Count                        int           `json:"count"`
		EventID                      int           `json:"eventID"`
		Point                        int           `json:"point"`
		AvaliableCount               int           `json:"avaliableCount"`
		CurrentTime                  int           `json:"currentTime"`
		Theme                        int           `json:"theme"`
		NewCreator                   interface{}   `json:"newCreator"`
		Token                        interface{}   `json:"token"`
		NameToken                    interface{}   `json:"nameToken"`
		RedEnvelopeID                string        `json:"redEnvelopeID"`
		CustomizedName               string        `json:"customizedName"`
		RecommendRoomID              string        `json:"recommendRoomID"`
		RecommendStreamerDisplayInfo interface{}   `json:"recommendStreamerDisplayInfo"`
		Type                         int           `json:"type"`
		InfoType                     int           `json:"infoType"`
		RedenvelopeDialogURL         string        `json:"redenvelopeDialogURL"`
		ButtonDecorationURL          string        `json:"buttonDecorationURL"`
		CountdownIconURL             string        `json:"countdownIconURL"`
		EntryIconURL                 string        `json:"entryIconURL"`
		CountBackgroundColor         string        `json:"countBackgroundColor"`
		ButtonTextColor              string        `json:"buttonTextColor"`
		AdditionalTextAndButtonColor string        `json:"additionalTextAndButtonColor"`
		IsBox                        bool          `json:"isBox"`
		NextEventType                int           `json:"nextEventType"`
		GiftIDs                      []interface{} `json:"giftIDs"`
		ToArchiveEventID             int           `json:"toArchiveEventID"`
	} `json:"redEnvelopeEventInfo"`
	RedEnvelopeEventEndMsg   interface{} `json:"redEnvelopeEventEndMsg"`
	RedEnvelopeRegionInfo    interface{} `json:"redEnvelopeRegionInfo"`
	VoteInfo                 interface{} `json:"voteInfo"`
	RockViewersMsg           interface{} `json:"rockViewersMsg"`
	MyArmyOverviewMsg        interface{} `json:"myArmyOverviewMsg"`
	EnterAnimationMsg        interface{} `json:"enterAnimationMsg"`
	DeathExemptionMedalInfo  interface{} `json:"deathExemptionMedalInfo"`
	TimeSyncMsg              interface{} `json:"timeSyncMsg"`
	ArmyInvitationMsg        interface{} `json:"armyInvitationMsg"`
	FreshUserEnterMsg        interface{} `json:"freshUserEnterMsg"`
	PmStreamerInfo           interface{} `json:"pmStreamerInfo"`
	RedEnvelopeRecommendMsg  interface{} `json:"redEnvelopeRecommendMsg"`
	MonsterInfo              interface{} `json:"monsterInfo"`
	ToastMsg                 interface{} `json:"toastMsg"`
	MarqueeMsg               interface{} `json:"marqueeMsg"`
	StreamerAssistantMsg     interface{} `json:"streamerAssistantMsg"`
	EventCommentMsg          interface{} `json:"eventCommentMsg"`
	LaborRewardRankup        interface{} `json:"laborRewardRankup"`
	DayChange                interface{} `json:"dayChange"`
	LaborReceiveRewardMsg    interface{} `json:"laborReceiveRewardMsg"`
	MissionRemind            interface{} `json:"missionRemind"`
	PkActionMsg              interface{} `json:"pkActionMsg"`
	PkScore                  interface{} `json:"pkScore"`
	AnimationInfo            interface{} `json:"animationInfo"`
	AsyncAnimationInfo       interface{} `json:"asyncAnimationInfo"`
	PromotionIAPInfo         interface{} `json:"promotionIAPInfo"`
	DailyTaskInfo            interface{} `json:"dailyTaskInfo"`
	RisingStarMsg            interface{} `json:"risingStarMsg"`
	GroupCallActionMsg       interface{} `json:"groupCallActionMsg"`
	GroupCallMemberStatusMsg interface{} `json:"groupCallMemberStatusMsg"`
	ChangeRoomMsg            interface{} `json:"changeRoomMsg"`
	EndStreamMsg             interface{} `json:"endStreamMsg"`
	UpdateFeatures           interface{} `json:"updateFeatures"`
	GiftBackupStatusMsg      interface{} `json:"giftBackupStatusMsg"`
	GiftBackupSendMsg        interface{} `json:"giftBackupSendMsg"`
	GiftBackupPayMsg         interface{} `json:"giftBackupPayMsg"`
	ArmySubscriptionMsg      interface{} `json:"armySubscriptionMsg"`
	SubscriptionMsg          interface{} `json:"subscriptionMsg"`
	PkRevenge                interface{} `json:"pkRevenge"`
	BoxgachaLevelUpMsg       interface{} `json:"boxgachaLevelUpMsg"`
	ArmyGiftMsg              interface{} `json:"armyGiftMsg"`
	LimitedGiftMsgList       interface{} `json:"limitedGiftMsgList"`
	GameInfoMsg              interface{} `json:"gameInfoMsg"`
	ShoppingCartMsg          interface{} `json:"shoppingCartMsg"`
	WarnStreamMsg            interface{} `json:"warnStreamMsg"`
	ClipOnLiveCreatedMsg     interface{} `json:"clipOnLiveCreatedMsg"`
	EpicMomentMsg            interface{} `json:"epicMomentMsg"`
	AiCohostMsg              interface{} `json:"aiCohostMsg"`
	AiCohostGameInfoMsg      interface{} `json:"aiCohostGameInfoMsg"`
	AiCohostGameResultMsg    interface{} `json:"aiCohostGameResultMsg"`
	Subtitle                 interface{} `json:"subtitle"`
	EngagementRewardMsg      interface{} `json:"engagementRewardMsg"`
	StarBoostRewardMsg       interface{} `json:"starBoostRewardMsg"`
}
