package wasm

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/Trey2k/OpenStreaming/webApp/app/database"
	"github.com/gorilla/mux"
)

func OverlayEditorHandler(w http.ResponseWriter, r *http.Request, id int) {
	vars := mux.Vars(r)

	usr := database.GetUserByID(id)

	p := page{
		Title:             "OpenStreaming - Overlay",
		Token:             vars["id"],
		DisplayName:       usr.HelixClient.UserData.DisplayName,
		ProfilePicture:    usr.HelixClient.UserData.ProfileImageURL,
		BackgroundPicture: usr.HelixClient.UserData.OfflineImageURL,
	}

	err := common.Templates.WasmPage.ExecuteTemplate(w, "editor", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing template:\n%s\n", err)
		return
	}

}
