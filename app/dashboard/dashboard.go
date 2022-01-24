package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/database"
)

type Page struct {
	Title        string
	LoggedIn     bool
	StringOne    string
	StringTwo    string
	CustomJS     bool
	CustomJSPath string
}

//GetHomePage reutnrs the home page
func GetHomePage(rw http.ResponseWriter, req *http.Request, id int) {
	usr := database.GetUser(id)

	p := Page{
		Title:        "OpenStreaming - Dashboard",
		StringOne:    usr.HelixClient.UserData.DisplayName,
		LoggedIn:     true,
		CustomJS:     true,
		CustomJSPath: "/static/js/dashboard.js",
	}

	err := common.Templates.DashboardPage.ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}

}
