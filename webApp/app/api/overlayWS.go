package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/webApp/app/twitch/eventSub"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/Trey2k/OpenStreaming/webApp/app/database"
	"github.com/gorilla/websocket"
)

type MessageType int

type MessageStruct struct {
	Type    MessageType
	Overlay *database.OverlayStruct
}

const (
	InvalidMessage = MessageType(iota)
	GetOverlay
	Return
)

var clients = make(map[*websocket.Conn]*database.UserStruct)

var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == os.Getenv("URL")
		},
	}
}

func messageHandler(overlay *database.OverlayStruct, msg *MessageStruct) error {
	switch msg.Type {
	case GetOverlay:
		toSend := &MessageStruct{
			Type:    Return,
			Overlay: overlay,
		}
		err := overlay.Websocket.WriteJSON(toSend)
		if err != nil {
			return err
		}
		return nil

	}
	return fmt.Errorf("unknown message type: %d", msg.Type)
}

func OverlayWSHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing form:\n%s\n", err)
		return
	}

	if len(r.Form["token"]) < 1 {
		w.WriteHeader(http.StatusForbidden)
		common.Loggers.Info.Printf("No token found for overlay request\n")
		return
	}
	key := r.Form["token"][0]

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while upgrading websocket connection:\n%s\n", err)
		return
	}

	fmt.Println(key)
	overlay, err := database.GetOverlayByKey(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while getting overlay from database:\n%s\n", err)
		return
	}

	user := database.GetUserByID(overlay.ID)
	user.Overlay.Websocket = ws
	clients[ws] = user

	eventSub.SubscribeTwitchEvents(user)

	ws.SetCloseHandler(func(code int, text string) error {
		clients[ws] = nil
		user.Overlay.Websocket = nil
		common.Loggers.Info.Printf("Dropped WS Connection with %d\n", user.ID)
		return nil
	})

	common.Loggers.Info.Printf("Opened WS Connection with %d\n", user.ID)

	for {
		msg := &MessageStruct{}
		err := ws.ReadJSON(msg)
		if err != nil {
			break
		}

		err = messageHandler(user.Overlay, msg)
		if err != nil {
			common.Loggers.Error.Printf("Error while processing ws message:\n%s\n", err)
		}

	}

}
