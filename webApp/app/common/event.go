package common

type EventStruct struct {
	Type EventType
	Data TwitchEventStruct
}

type EventType int

const (
	InvalidEvent = EventType(iota)
	TestEvent
	TwitchMessageEvent
	TwitchFollow
)

type TwitchEventStruct struct {
	Channel        string
	DisplayName    string
	ProfilePicture string
	UserID         string
	MessageContent string
}
