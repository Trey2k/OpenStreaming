package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

type AuthenticatedHandlerFunc func(w http.ResponseWriter, r *http.Request, id int)

func AuthenticatedMW(handler AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		IsAuthenticated, id := IsAuthenticated(w, r)
		if !IsAuthenticated {
			http.Redirect(w, r, "/login", http.StatusFound)
			common.Loggers.Info.Printf("Unauthenticated request to %s ip: %s\n", r.URL, common.GetIP(r))
			return
		}

		handler(w, r, id)
	}
}
