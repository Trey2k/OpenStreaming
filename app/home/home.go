package home

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/twitchHelix"
	"github.com/Trey2k/OpenStreaming/app/user"
)

//GetHomePage reutnrs the home page
func GetHomePage(rw http.ResponseWriter, req *http.Request) {

	isAuthenticated, id := isAuthenticated(rw, req)
	fmt.Println("Testing ", id)
	if !isAuthenticated {
		type Page struct {
			Title         string
			TwitchAuthURL string
		}

		p := Page{
			Title: "OpenStreaming - Login",
			TwitchAuthURL: fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
				os.Getenv("TwitchClientID"), "https://weaselfoss.dev/twitch", twitchHelix.Scope),
		}

		err := common.Templates.HomeTemplates["login"].ExecuteTemplate(rw, "base", p)
		if err != nil {
			panic(err)
		}
		return
	}

	usr, err := user.GetUser(int(id))
	if err != nil {
		panic(err)
	}

	type Page struct {
		Title       string
		DisplayName string
	}

	p := Page{
		Title:       "OpenStreaming - Home",
		DisplayName: usr.HelixClient.User.DisplayName,
	}

	err = common.Templates.HomeTemplates["home"].ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}

}
