package dashboard

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/Trey2k/OpenStreaming/app/twitch/helix"
)

func GetLoginPage(w http.ResponseWriter, r *http.Request) {

	p := page{
		Title: "OpenStreaming - Login",
		StringOne: fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
			os.Getenv("TwitchClientID"), fmt.Sprintf("%s/twitch", os.Getenv("URL")), helix.Scope),
		LoggedIn: false,
	}

	err := common.Templates.LoginPage.ExecuteTemplate(w, "base", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.Loggers.Error.Printf("Error while parsing template:\n%s\n", err)
		return
	}

}
