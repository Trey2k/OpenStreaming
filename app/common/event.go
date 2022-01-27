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
	TwitchEventSub
)

type TwitchMessageEventStruct struct {
	Channel        string
	UserDisplay    string
	UserID         string
	MessageContent string
}
