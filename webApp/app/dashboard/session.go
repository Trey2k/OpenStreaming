package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/Trey2k/OpenStreaming/webApp/app/database"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge: 60 * 15,
		Secure: true,
	}

}

func createSession(w http.ResponseWriter, r *http.Request, id int) error {
	session, _ := store.Get(r, "session-token")
	session.Values["id"] = id
	err := session.Save(r, w)
	return err
}

func getSession(s *sessions.Session) (int, bool) {
	val := s.Values["id"]
	id, ok := val.(int)
	if !ok {
		return id, false
	}
	return id, true
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) (bool, int) {
	session, err := store.Get(r, "session-token")
	if err == nil {
		id, authenitcated := getSession(session)
		if authenitcated {
			return true, id
		}
	}

	return false, 0
}

func TwitchOAuthEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		codes, ok := r.URL.Query()["code"]
		if ok && len(codes) == 0 {
			http.Redirect(w, r, "/login", http.StatusFound)
			common.Loggers.Info.Printf("No twitch token given during login. ip: %s\n", common.GetIP(r))
			return
		}

		user, err := database.NewUser(codes[0])
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			common.Loggers.Info.Printf("Invalid twitch token given during login. ip: %s\n", common.GetIP(r))
			return
		}

		err = createSession(w, r, user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			common.Loggers.Error.Printf("Error while creating session:\n%s\n", err)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusFound)

	}
}
