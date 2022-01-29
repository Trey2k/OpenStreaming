package common

type EventStruct struct {
	Type EventType
	Data interface{}
}

type EventType int

const (
	InvalidEvent = EventType(iota)
	TestEvent
	TwitchMessageEvent
	TwitchFollow
)

type TwitchMessageEventStruct struct {
	Channel        string
	UserDisplay    string
	UserID         string
	MessageContent string
}

type TwitchFollowEventStruct struct {
	DisplayName    string
	ProfilePicture string
}
