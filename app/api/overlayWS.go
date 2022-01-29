package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/twitch/eventSub"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/database"
	"github.com/gorilla/websocket"
)

type MessageType string

type MessageStruct struct {
	Type    MessageType
	Index   int
	Overlay *database.OverlayStruct
	Modeule *database.OverlayModule
}

const (
	InvalidMessage = ""
	GetOverlay     = "getOverlay"
	Return         = "return"
	UpdateModeul   = "updateModule"
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
	overlay = user.Overlay
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

		switch msg.Type {
		case GetOverlay:
			toSend := &MessageStruct{
				Type:    Return,
				Overlay: overlay,
			}
			err = ws.WriteJSON(toSend)
			if err != nil {
				common.Loggers.Error.Printf("Error while writing overlay to websocket:\n%s\n", err)
			}
		}

	}

}
