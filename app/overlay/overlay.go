package overlay

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

type page struct {
	Title string
}

func OverlayHandler(w http.ResponseWriter, r *http.Request) {

	p := page{
		Title: "OpenStreaming - Overlay",
	}

	err := common.Templates.OverlayPage.ExecuteTemplate(w, "base", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing template:\n%s\n", err)
		return
	}

}
