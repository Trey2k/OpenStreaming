package api

import (
	"encoding/json"
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/dashboard"
	"github.com/Trey2k/OpenStreaming/app/database"
)

func GetEventHandler(rw http.ResponseWriter, req *http.Request) {
	isAuthenticate, id := dashboard.IsAuthenticated(rw, req)
	if !isAuthenticate {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	usr := database.GetUser(id)
	events := usr.GetEvents()

	err := json.NewEncoder(rw).Encode(events)
	if err != nil {
		panic(err)
	}
}
