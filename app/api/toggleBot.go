package api

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/dashboard"
	"github.com/Trey2k/OpenStreaming/app/database"
)

func ToggleBotHandler(rw http.ResponseWriter, req *http.Request) {
	isAuthenticate, id := dashboard.IsAuthenticated(rw, req)
	if !isAuthenticate {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	usr := database.GetUser(id)
	if usr.HelixClient.ChatBot.Status {
		usr.HelixClient.ChatBot.Stop()
	} else {
		usr.HelixClient.ChatBot.Start()
	}
}
