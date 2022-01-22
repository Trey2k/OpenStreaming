package home

import (
	"fmt"
	"net/http"

	"github.com/Trey2k/OpenStreaming/app/user"
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

func createSession(req *http.Request, rw http.ResponseWriter, userID int64) {
	session, err := store.Get(req, "session-token")
	session.Values["id"] = userID
	err = session.Save(req, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getSession(s *sessions.Session) (int64, bool) {
	val := s.Values["id"]
	id, ok := val.(int64)
	if !ok {
		return id, false
	}
	return id, true
}

func isAuthenticated(rw http.ResponseWriter, req *http.Request) (bool, int64) {
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
		fmt.Println(codes[0])
		if ok && len(codes) > 0 {
			user, err := user.NewUser(codes[0])
			if err != nil {
				panic(err)
			}

			createSession(req, rw, int64(user.ID))

			http.Redirect(rw, req, "https://weaselfoss.dev/", 200)
		}
	}
}
