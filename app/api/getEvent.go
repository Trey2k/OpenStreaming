package api

import (
	"encoding/json"
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/dashboard"
	"github.com/Trey2k/OpenStreaming/app/database"
)

func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticate, id := dashboard.IsAuthenticated(w, r)
	if !isAuthenticate {
		w.WriteHeader(http.StatusForbidden)
		common.Loggers.Info.Printf("Unauthenticated request to %s ip: %s\n", r.URL, common.GetIP(r))
		return
	}
	usr := database.GetUserByID(id)
	events := usr.GetEvents()

	err := json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while encoding events:\n%s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
