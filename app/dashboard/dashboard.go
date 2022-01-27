package dashboard

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/database"
)

type page struct {
	Title             string
	LoggedIn          bool
	DisplayName       string
	ProfilePicture    string
	BackgroundPicture string
	OverlayURL        string
	StringOne         string
}

//GetHomePage reutnrs the home page
func GetHomePage(w http.ResponseWriter, r *http.Request, id int) {
	usr := database.GetUser(id)

	p := page{
		Title:             "OpenStreaming - Dashboard",
		DisplayName:       usr.HelixClient.UserData.DisplayName,
		ProfilePicture:    usr.HelixClient.UserData.ProfileImageURL,
		BackgroundPicture: usr.HelixClient.UserData.OfflineImageURL,
		OverlayURL:        fmt.Sprintf("%s/overlay?id=%s", os.Getenv("URL"), usr.Overlay.Key),
		LoggedIn:          true,
	}

	err := common.Templates.DashboardPage.ExecuteTemplate(w, "base", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing template:\n%s\n", err)
		return
	}

}
