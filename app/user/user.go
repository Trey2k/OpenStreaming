package user

import (
	"context"

	"github.com/Trey2k/OpenStreaming/app/database"
	"github.com/Trey2k/OpenStreaming/app/twitchHelix"
)

type UserStruct struct {
	ID          int
	TwitchID    string
	DiscordID   string
	HelixClient *twitchHelix.HelixClientStruct
}

func NewUser(refreshToken string) (*UserStruct, error) {
	user := &UserStruct{}
	var err error
	user.HelixClient, err = twitchHelix.NewHelixClient(refreshToken)
	err = user.FetchDB()
	if err != nil {
		err = user.CreateUser()
		if err != nil {
			return nil, err
		}
		err = user.FetchDB()
	}
	return user, err
}

func (user *UserStruct) FetchDB() error {
	conn, err := database.ConnectDB()
	if err != nil {
		return err
	}

	err = conn.QueryRow(context.Background(), `SELECT * FROM public.users WHERE "twitchID"=$1`, user.HelixClient.User.ID).Scan(&user.ID, &user.TwitchID, &user.DiscordID)
	return err
}

func (user *UserStruct) CreateUser() error {
	conn, err := database.ConnectDB()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), `INSERT INTO public.users("twitchID") VALUES ($1) `, user.HelixClient.User.ID)
	user.TwitchID = user.HelixClient.User.ID
	if err != nil {
		panic(err)
	}
	return err
}
