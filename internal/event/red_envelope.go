package event

// RedEnvelope represents a red envelope packet
type RedEnvelopeInfo struct {
	ID        string
	RoomID    string
	Count     int
	StartTime int
	EndTime   int
}

func (r *RedEnvelopeInfo) Type() string {
	return "red_envelope"
}
