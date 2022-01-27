package api

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/dashboard"
	"github.com/Trey2k/OpenStreaming/app/database"
)

func ToggleBotHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticate, id := dashboard.IsAuthenticated(w, r)
	if !isAuthenticate {
		w.WriteHeader(http.StatusForbidden)
		common.Loggers.Info.Printf("Unauthenticated request to %s ip: %s\n", r.URL, common.GetIP(r))
		return
	}
	usr := database.GetUserByID(id)
	w.WriteHeader(http.StatusOK)
	if usr.HelixClient.ChatBot.Status {
		usr.HelixClient.ChatBot.Stop()
		return
	}
	usr.HelixClient.ChatBot.Start()

}
