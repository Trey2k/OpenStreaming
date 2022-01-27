package eventSub

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/database"
	"github.com/Trey2k/OpenStreaming/app/twitch/helix"
)

func EventSubHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		eventType := req.Header.Get("twitch-eventsub-message-type")
		switch eventType {
		case "webhook_callback_verification":
			println("sub successful")
			challange := &common.ChallangeData{}
			json.NewDecoder(req.Body).Decode(challange)
			rw.Write([]byte(challange.Challenge))
		case "notification":
			event := &common.EventSubData{}
			json.NewDecoder(req.Body).Decode(event)

			var user = database.GetUserByTwitchID(event.Event.BroadcasterUserID)

			err := user.FetchDB()
			if err != nil {
				panic(err)
			}

			data := &common.EventStruct{
				Type: common.TwitchEventSub,
				Data: event,
			}

			user.SendEvent(data)
		}
	}
}

func SubscribeTwitchEvents(user *database.UserStruct) {
	data := &common.SubscribeData{
		Type:    "channel.follow", //TODO: add channel.subscribe if the user is affiliate
		Version: "1",
		Condition: common.ConditionData{
			BroadcasterUserID: user.HelixClient.UserData.ID,
		},
		Transport: common.TransportData{
			Method:   "webhook",
			Callback: os.Getenv("URL") + "/eventsub",
			Secret:   os.Getenv("TwitchClientSecret"),
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "https://api.twitch.tv/helix/eventsub/subscriptions", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")
	resp, err := helix.DoAppRequest(request)
	if err != nil {
		panic(err)
	}

	user.EventSubData = &common.EventSubData{}

	err = json.NewDecoder(resp.Body).Decode(user.EventSubData)
	if err != nil {
		panic(err)
	}
}
