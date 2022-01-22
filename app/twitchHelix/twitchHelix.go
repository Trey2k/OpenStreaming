package twitchHelix

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const Scope = "channel:moderate chat:edit chat:read whispers:read whispers:edit user:read:follows user:read:subscriptions user:read:email user:read:broadcast user:read:blocked_users user:manage:blocked_users user:edit:follows user:edit moderator:manage:chat_settings moderator:read:chat_settings moderator:manage:automod_settings moderator:read:automod_settings moderator:manage:automod moderator:manage:blocked_terms moderator:read:blocked_terms moderator:manage:banned_users moderation:read clips:edit channel:read:subscriptions channel:read:stream_key channel:read:redemptions channel:read:predictions channel:read:polls channel:read:hype_train channel:read:goals channel:read:editors channel:manage:videos channel:manage:schedule channel:manage:redemptions channel:manage:predictions channel:manage:polls channel:manage:extensions channel:manage:broadcast channel:edit:commercial bits:read analytics:read:games analytics:read:extensions"

var AppToken TwitchRefresh

type HelixClientStruct struct {
	Refresh TwitchRefresh
	User    TwitchUserData
}

func init() {
	refreshAppToken()
}

func refreshAppToken() {
	data := url.Values{
		"client_id":     {os.Getenv("TwitchClientID")},
		"client_secret": {os.Getenv("TwitchClientSecret")},
		"grant_type":    {"client_credentials"},
		"scope":         {Scope},
	}
	resp, err := http.PostForm("https://id.twitch.tv/oauth2/token", data)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&AppToken)

	timer := time.NewTimer(time.Duration(AppToken.ExpiresIn))
	go func() {
		<-timer.C
		refreshAppToken()
	}()
}

func NewHelixClient(RefreshToken string) (*HelixClientStruct, error) {
	client := &HelixClientStruct{Refresh: TwitchRefresh{RefreshToken: RefreshToken}}

	err := client.refreshToken()
	if err != nil {
		return nil, err
	}

	err = client.getUserData()
	return client, err
}

func (client *HelixClientStruct) refreshToken() error {
	data := url.Values{
		"client_id":     {os.Getenv("TwitchClientID")},
		"client_secret": {os.Getenv("TwitchClientSecret")},
		"code":          {client.Refresh.RefreshToken},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {fmt.Sprintf("%s/twitch", os.Getenv("URL"))},
	}
	resp, err := http.PostForm("https://id.twitch.tv/oauth2/token", data)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(&client.Refresh)
}

func (client *HelixClientStruct) getUserData() error {
	request, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	if err != nil {
		return err
	}
	fmt.Println(client.Refresh.AccessToken)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Refresh.AccessToken))
	request.Header.Add("Client-Id", os.Getenv("TwitchClientID"))

	webClient := &http.Client{}

	fmt.Println(request)

	resp, err := webClient.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(&client.User)
}
