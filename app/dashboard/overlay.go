package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

func OverlayHandler(w http.ResponseWriter, r *http.Request, id int) {

	p := Page{
		Title: "OpenStreaming - Overlay",
	}

	err := common.Templates.OverlayPage.ExecuteTemplate(w, "min", p)
	if err != nil {
		panic(err)
	}

}
