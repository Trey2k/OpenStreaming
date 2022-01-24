package dashboard

import (
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/database"
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

func createSession(req *http.Request, rw http.ResponseWriter, id int) {
	session, err := store.Get(req, "session-token")
	session.Values["id"] = id
	err = session.Save(req, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getSession(s *sessions.Session) (int, bool) {
	val := s.Values["id"]
	id, ok := val.(int)
	if !ok {
		return id, false
	}
	return id, true
}

func IsAuthenticated(rw http.ResponseWriter, req *http.Request) (bool, int) {
	session, err := store.Get(req, "session-token")
	if err == nil {
		id, authenitcated := getSession(session)
		if authenitcated {
			return true, id
		}
	}

	return false, 0
}

func TwitchOAuthEndpoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		codes, ok := req.URL.Query()["code"]
		if ok && len(codes) > 0 {
			user, err := database.NewUser(codes[0])
			if err != nil {
				panic(err)
			}

			createSession(req, rw, user.ID)

			http.Redirect(rw, req, "https://weaselfoss.dev/", 200)
		}
	}
}
