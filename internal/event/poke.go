package event

type User struct {
	UserID      string `json:"userID"`
	DisplayName string `json:"displayName"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	OpenID      string `json:"openID"`
	Region      string `json:"region"`
}

type Poke struct {
	RoomID          string
	Sender          User
	Receiver        User
	IsPokeBack      bool
	CoolDownEndTime int
}

func (p *Poke) Type() string {
	return "poke"
}
