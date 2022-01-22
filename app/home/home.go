package home

import (
	"fmt"
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/common"
)

//GetHomePage reutnrs the home page
func GetHomePage(rw http.ResponseWriter, req *http.Request) {

	type Page struct {
		Title         string
		TwitchAuthURL string
		DisplayName   string
	}

	isAuthenticated, usr := isAuthenticated(rw, req)
	fmt.Println("Testing ", usr, isAuthenticated)
	if !isAuthenticated {
		http.Redirect(rw, req, "/login", 403)
		return
	}
	fmt.Println("test1")

	p := Page{
		Title:       "OpenStreaming - Home",
		DisplayName: usr.HelixClient.User.DisplayName,
	}

	err := common.Templates.HomeTemplates["home"].ExecuteTemplate(rw, "base", p)
	fmt.Println()
	if err != nil {
		panic(err)
	}

}
