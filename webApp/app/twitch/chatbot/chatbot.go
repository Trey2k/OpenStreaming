package chatbot

import (
	"fmt"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/gempir/go-twitch-irc/v2"
)

type ChatBot struct {
	Status    bool
	irc       *twitch.Client
	username  string
	eventChan chan *common.EventStruct
}

func NewChatBot(username string, token string, eventChan chan *common.EventStruct) (*ChatBot, error) {

	bot := &ChatBot{}

	bot.irc = twitch.NewClient(username, fmt.Sprintf("oauth:%s", token))
	bot.username = username
	bot.eventChan = eventChan

	bot.irc.OnPrivateMessage(bot.onMessage)

	return bot, nil
}

func (bot *ChatBot) UpdateToken(token string) {
	bot.irc.SetIRCToken(fmt.Sprintf("oauth:%s", token))
}

func (bot *ChatBot) onMessage(msg twitch.PrivateMessage) {

	event := common.EventStruct{
		Type: common.TwitchMessageEvent,
		Data: common.TwitchEventStruct{
			Channel:        msg.Channel,
			DisplayName:    msg.User.DisplayName,
			UserID:         msg.User.ID,
			MessageContent: msg.Message,
		},
	}
	bot.eventChan <- &event
}

func (bot *ChatBot) Sayf(message string, args ...interface{}) {
	bot.irc.Say(bot.username, fmt.Sprintf(message, args...))
}

func (bot *ChatBot) Start() {
	bot.irc.Join(bot.username)
	bot.Status = true
	go func() {
		bot.irc.Connect()
	}()
}

func (bot *ChatBot) Stop() {
	bot.Status = false
	bot.irc.Disconnect()
}
