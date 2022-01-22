package home

import (
	"encoding/gob"
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

	gob.Register(&user.UserStruct{})
}

func createSession(req *http.Request, rw http.ResponseWriter, user *user.UserStruct) {
	session, err := store.Get(req, "session-token")
	session.Values["user"] = user
	err = session.Save(req, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getSession(s *sessions.Session) (*user.UserStruct, bool) {
	val := s.Values["user"]
	user, ok := val.(*user.UserStruct)
	if !ok {
		return user, false
	}
	return user, true
}

func isAuthenticated(rw http.ResponseWriter, req *http.Request) (bool, *user.UserStruct) {
	session, err := store.Get(req, "session-token")
	if err == nil {
		user, authenitcated := getSession(session)
		if authenitcated {
			return true, user
		}
	}

	return false, nil
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

			createSession(req, rw, user)

			http.Redirect(rw, req, "https://weaselfoss.dev/", 200)
		}
	}
}
