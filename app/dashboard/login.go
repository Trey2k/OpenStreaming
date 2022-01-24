package dashboard

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/twitch/helix"
)

func GetLoginPage(rw http.ResponseWriter, req *http.Request) {

	p := Page{
		Title: "OpenStreaming - Login",
		StringOne: fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
			os.Getenv("TwitchClientID"), "https://weaselfoss.dev/twitch", helix.Scope),
		LoggedIn: false,
	}

	err := common.Templates.LoginPage.ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}

}
