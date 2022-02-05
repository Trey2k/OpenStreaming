package common

type EventStruct struct {
	Type    EventType
	Overlay interface{}
	WSid    int
	Data    TwitchEventStruct
}

type EventType int

const (
	InvalidEvent = EventType(iota)
	TestEvent
	TwitchMessageEvent
	TwitchFollow
	GetOverlay
	OverlayInfo
	SaveOverlay
)

type TwitchEventStruct struct {
	Channel        string
	DisplayName    string
	ProfilePicture string
	UserID         string
	MessageContent string
}
