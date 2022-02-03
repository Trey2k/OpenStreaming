package overlay

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/gorilla/mux"
)

type page struct {
	Title             string
	Token             string
	LoggedIn          bool
	DisplayName       string
	ProfilePicture    string
	BackgroundPicture string
}

func OverlayHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	p := page{
		Title: "OpenStreaming - Overlay",
		Token: vars["id"],
	}

	err := common.Templates.OverlayPage.ExecuteTemplate(w, "overlay", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing template:\n%s\n", err)
		return
	}

}
