package dashboard

import "net/http"

type AuthenticatedHandlerFunc func(rw http.ResponseWriter, r *http.Request, id int)

func AuthenticatedMW(handler AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		IsAuthenticated, id := IsAuthenticated(rw, req)
		if !IsAuthenticated {
			http.Redirect(rw, req, "/login", 403)
			return
		}

		handler(rw, req, id)
	}
}
