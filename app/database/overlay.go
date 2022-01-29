package database

import (
	"context"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type OverlayStruct struct {
	ID        int
	UserID    int
	Key       string
	Modules   map[int]*OverlayModule
	Websocket *websocket.Conn
}

func GetOverlayByUserID(userID int) (*OverlayStruct, error) {
	conn, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	toReturn := &OverlayStruct{}
	toReturn.UserID = userID
	err = conn.QueryRow(context.Background(), `SELECT id, "key" FROM public.overlays WHERE "userID"=$1`, userID).Scan(&toReturn.ID, &toReturn.Key)
	if err != nil {
		return nil, err
	}

	err = toReturn.GetModules()
	return toReturn, err
}

func GetOverlayByKey(key string) (*OverlayStruct, error) {
	conn, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	toReturn := &OverlayStruct{}
	toReturn.Key = key
	err = conn.QueryRow(context.Background(), `SELECT id, "userID" FROM public.overlays WHERE "key"=$1`, key).Scan(&toReturn.ID, &toReturn.UserID)
	if err != nil {
		return nil, err
	}

	err = toReturn.GetModules()
	return toReturn, err
}

func createOverlay(user *UserStruct) error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	key := uuid.NewV4().String()

	user.Overlay = &OverlayStruct{
		UserID: user.ID,
		Key:    key,
	}

	err = conn.QueryRow(context.Background(), `INSERT INTO public.overlays("userID", "key") VALUES ($1, $2) RETURNING "id";`, user.ID, key).Scan(&user.Overlay.ID)
	if err != nil {
		return err
	}

	err = user.Overlay.GetModules()
	return err
}
