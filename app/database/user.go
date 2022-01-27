package database

import (
	"context"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/twitch/helix"
)

type UserStruct struct {
	ID        int
	TwitchID  string
	DiscordID string

	Events      []common.EventStruct
	HelixClient *helix.HelixClientStruct
	Overlay     *OverlayStruct
	eventChan   chan common.EventStruct
}

var (
	users map[int]*UserStruct
)

func init() {
	users = make(map[int]*UserStruct)
}

func NewUser(refreshToken string) (*UserStruct, error) {
	user := &UserStruct{}

	var err error
	user.eventChan = make(chan common.EventStruct)
	user.HelixClient, err = helix.NewHelixClient(refreshToken, UpdateRefreshToken, user.eventChan)
	err = user.FetchDB()
	if err != nil {
		err = user.CreateUser()
		if err != nil {
			return nil, err
		}
		err = user.FetchDB()
		if err != nil {
			return nil, err
		}
	}
	users[user.ID] = user

	user.Overlay, err = GetOverlayByUserID(user.ID)
	if err != nil {
		err = createOverlay(user)
		if err != nil {
			return user, err
		}
	}

	go user.ListenForEvents()

	return user, err
}

func (user *UserStruct) ListenForEvents() {
	for {
		select {
		case event := <-user.eventChan:
			user.SendEvent(event)
		}
	}
}

func GetUser(id int) *UserStruct {
	return users[id]
}

func (user *UserStruct) FetchDB() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	err = conn.QueryRow(context.Background(), `SELECT id, "twitchID", "discordID" FROM public.users WHERE "twitchID"=$1`, user.HelixClient.UserData.ID).Scan(&user.ID, &user.TwitchID, &user.DiscordID)
	return err
}

func (user *UserStruct) SendEvent(event common.EventStruct) error {
	user.Events = append(user.Events, event)
	if user.Overlay.Websocket != nil {
		err := user.Overlay.Websocket.WriteJSON(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func (user *UserStruct) GetEvents() []common.EventStruct {
	temp := user.Events
	return temp
}

func UpdateRefreshToken(token, twitchID string) error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), `UPDATE public.users SET "refreshToken"=$1 WHERE "twitchID"=$2;`, token, twitchID)
	return err
}

func (user *UserStruct) CreateUser() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), `INSERT INTO public.users("twitchID", "refreshToken") VALUES ($1, $2) `, user.HelixClient.UserData.ID, user.HelixClient.Refresh.RefreshToken)
	user.TwitchID = user.HelixClient.UserData.ID
	return err
}
