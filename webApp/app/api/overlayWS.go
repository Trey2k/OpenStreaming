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

// This exists as a hack to let json parse overlay nicly
type ReciviedEventStruct struct {
	Type    common.EventType
	Overlay *database.OverlayStruct
	wsID    int
	Data    common.TwitchEventStruct
}

type EventType int

type TwitchEventStruct struct {
	Channel        string
	DisplayName    string
	ProfilePicture string
	UserID         string
	MessageContent string
}

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

func saveOverlay(msg *ReciviedEventStruct, overlay *database.OverlayStruct) error {
	var newIDs []int

	for k, v := range msg.Overlay.ModuleInfo {
		if _, ok := overlay.ModuleInfo[k]; ok && !v.IsNew {
			overlay.ModuleInfo[k].Update(v)
			continue
		}

		id, err := overlay.NewModule(v)
		if err != nil {
			return err
		}

		newIDs = append(newIDs, id)

	}

	for k, v := range overlay.ModuleInfo {
		if _, ok := msg.Overlay.ModuleInfo[k]; !ok && !common.ContainsInt(newIDs, k) {
			v.Delete()
			delete(overlay.ModuleInfo, k)
			continue
		}
		v.IsNew = false
	}

	toSend := &common.EventStruct{
		Type:    common.OverlayInfo,
		Overlay: overlay,
	}

	for _, ws := range overlay.Websockets {
		err := ws.WriteJSON(toSend)
		if err != nil {
			return err
		}
	}
	return nil

}

func messageHandler(overlay *database.OverlayStruct, msg *ReciviedEventStruct) error {
	switch msg.Type {
	case common.GetOverlay:
		toSend := &common.EventStruct{
			Type:    common.OverlayInfo,
			Overlay: overlay,
		}
		err := overlay.Websockets[msg.wsID].WriteJSON(toSend)
		if err != nil {
			return err
		}
		return nil
	case common.SaveOverlay:
		return saveOverlay(msg, overlay)
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

	overlay, err := database.GetOverlayByKey(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while getting overlay from database:\n%s\n", err)
		return
	}

	user := database.GetUserByID(overlay.ID)
	var wsID = len(user.Overlay.Websockets)
	if user.Overlay.Websockets == nil {
		user.Overlay.Websockets = make(map[int]*websocket.Conn)
	}

	user.Overlay.Websockets[wsID] = ws

	clients[ws] = user

	eventSub.SubscribeTwitchEvents(user)

	ws.SetCloseHandler(func(code int, text string) error {
		delete(clients, ws)
		delete(user.Overlay.Websockets, wsID)

		common.Loggers.Info.Printf("Dropped WS Connection with %d\n", user.ID)
		return nil
	})

	common.Loggers.Info.Printf("Opened WS Connection with %d\n", user.ID)

	for {
		msg := &ReciviedEventStruct{
			wsID: wsID,
		}
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
