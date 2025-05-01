package event

type UserJoined struct {
	RoomID   string
	UserID   string
	Username string
	Picture  string
}

func (u *UserJoined) Type() string {
	return "user_joined"
}
