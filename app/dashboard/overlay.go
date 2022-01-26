package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

func OverlayHandler(w http.ResponseWriter, r *http.Request, id int) {

	p := Page{
		Title:        "OpenStreaming - Overlay",
		CustomJS:     true,
		CustomJSPath: "/static/js/overlay.js",
	}

	err := common.Templates.OverlayPage.ExecuteTemplate(w, "min", p)
	if err != nil {
		panic(err)
	}

}
