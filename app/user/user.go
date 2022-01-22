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

func GetUser(id int) (*UserStruct, error) {
	user := &UserStruct{}

	conn, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	var oldToken string
	err = conn.QueryRow(context.Background(), `SELECT * FROM public.users WHERE id=$1`, id).Scan(&user.ID, &user.TwitchID, &user.DiscordID, &oldToken)
	if err != nil {
		return nil, err
	}

	user.HelixClient, err = twitchHelix.NewHelixClient(oldToken)

	return user, err
}

func (user *UserStruct) FetchDB() error {
	conn, err := database.ConnectDB()
	if err != nil {
		return err
	}

	var oldToken string
	err = conn.QueryRow(context.Background(), `SELECT * FROM public.users WHERE "twitchID"=$1`, user.HelixClient.User.ID).Scan(&user.ID, &user.TwitchID, &user.DiscordID, &oldToken)
	return err
}

func (user *UserStruct) CreateUser() error {
	conn, err := database.ConnectDB()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), `INSERT INTO public.users("twitchID", "refreshToken") VALUES ($1, $2) `, user.HelixClient.User.ID, user.HelixClient.Refresh.RefreshToken)
	user.TwitchID = user.HelixClient.User.ID
	return err
}
