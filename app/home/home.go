package home

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

//GetHomePage reutnrs the home page
func GetHomePage(rw http.ResponseWriter, req *http.Request) {

	type Page struct {
		Title       string
		LoggedIn    bool
		DisplayName string
	}

	isAuthenticated, usr := isAuthenticated(rw, req)

	if !isAuthenticated {
		http.Redirect(rw, req, "/login", 403)
		return
	}

	p := Page{
		Title:       "OpenStreaming - Home",
		DisplayName: usr.HelixClient.User.DisplayName,
		LoggedIn:    true,
	}

	err := common.Templates.HomePage.ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}

}
