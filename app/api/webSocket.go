package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Msg string
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{
		
	}
	upgrader.CheckOrigin = checkOrgin
}

func checkOrgin(r *http.Request) bool {
	return true
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true

	fmt.Println("Client connected!")
}

func Echo() {
	for {
		<-broadcast
		latlong := "Hello Client!"
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
