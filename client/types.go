package client

import "github.com/gorilla/websocket"

// Client represents a WebSocket client with a message handler
type Client struct {
	conn      *websocket.Conn
	channels  []int
	onMessage func(msg Message)
}

type ClientConfig struct {
	// Identity struct {
	// 	Username string
	// 	Password string
	// }
	Channels []int
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
	Channel  string
	UserID   string
	Username string
	Picture  string
	Content  string
	Tags     []string
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
	SubscriberEnterMsg       interface{} `json:"subscriberEnterMsg"`
	ReactMsg                 interface{} `json:"reactMsg"`
	DisplayUserMsg           interface{} `json:"displayUserMsg"`
	GuardianInfo             interface{} `json:"guardianInfo"`
	GiftMsg                  interface{} `json:"giftMsg"`
	GlobalLuckyBagMsg        interface{} `json:"globalLuckyBagMsg"`
	LiveStreamBroadcastMsg   interface{} `json:"liveStreamBroadcastMsg"`
	GlobalInfoUpdateMsg      interface{} `json:"globalInfoUpdateMsg"`
	GameMsg                  interface{} `json:"gameMsg"`
	Liveinfo                 interface{} `json:"liveinfo"`
	LiveinfoChange           interface{} `json:"liveinfoChange"`
	PokeInfo                 interface{} `json:"pokeInfo"`
	PollInfo                 interface{} `json:"pollInfo"`
	MediaMessage             interface{} `json:"mediaMessage"`
	GlobalAnnouncementMsg    interface{} `json:"globalAnnouncementMsg"`
	LiveStreamPromotionMsg   interface{} `json:"liveStreamPromotionMsg"`
	TriviaMsg                interface{} `json:"triviaMsg"`
	RedEnvelopeEventInfo     interface{} `json:"redEnvelopeEventInfo"`
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
