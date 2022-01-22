package home

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/twitchHelix"
)

func GetLoginPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title         string
		TwitchAuthURL string
		DisplayName   string
	}

	fmt.Println("test")
	p := Page{
		Title: "OpenStreaming - Login",
		TwitchAuthURL: fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
			os.Getenv("TwitchClientID"), "https://weaselfoss.dev/twitch", twitchHelix.Scope),
	}

	err := common.Templates.HomeTemplates["login"].ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}

}
