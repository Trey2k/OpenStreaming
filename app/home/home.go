package home

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
)

const scope = "channel:moderate chat:edit chat:read whispers:read whispers:edit user:read:follows user:read:subscriptions user:read:email user:read:broadcast user:read:blocked_users user:manage:blocked_users user:edit:follows user:edit moderator:manage:chat_settings moderator:read:chat_settings moderator:manage:automod_settings moderator:read:automod_settings moderator:manage:automod moderator:manage:blocked_terms moderator:read:blocked_terms moderator:manage:banned_users moderation:read clips:edit channel:read:subscriptions channel:read:stream_key channel:read:redemptions channel:read:predictions channel:read:polls channel:read:hype_train channel:read:goals channel:read:editors channel:manage:videos channel:manage:schedule channel:manage:redemptions channel:manage:predictions channel:manage:polls channel:manage:extensions channel:manage:broadcast channel:edit:commercial bits:read analytics:read:games analytics:read:extensions"

//GetHomePage reutnrs the home page
func GetHomePage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title         string
		TwitchAuthURL string
	}

	p := Page{
		Title: "OpenStreaming - Home",
		TwitchAuthURL: fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
			os.Getenv("TwitchClientID"), "https://weaselfoss.dev/twitch", scope),
	}

	err := common.Templates.HomeTemplates["home"].ExecuteTemplate(rw, "base", p)
	if err != nil {
		panic(err)
	}
}
