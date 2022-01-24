package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Trey2k/OpenStreaming/app/api"
	"github.com/Trey2k/OpenStreaming/app/dashboard"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.HandleFunc("/", httpInterceptor(router))
	router.Handle("/", http.RedirectHandler("/home", 200)).Methods("GET")

	// Main site
	router.HandleFunc("/home", dashboard.AuthenticatedMW(dashboard.GetHomePage)).Methods("GET")
	router.HandleFunc("/login", dashboard.GetLoginPage).Methods("GET")
	router.HandleFunc("/overlay", dashboard.AuthenticatedMW(dashboard.OverlayHandler)).Methods("GET")
	router.HandleFunc("/twitch", dashboard.TwitchOAuthEndpoint()).Methods("GET")

	// Api endpoints
	router.HandleFunc("/api/getEvents", api.GetEventHandler).Methods("GET")
	router.HandleFunc("/api/toggleBot", api.ToggleBotHandler).Methods("GET")
	// Static file server
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("/root/resources/static")))
	http.Handle("/static/", fileServer)

	server := &http.Server{
		Addr: ":443",
	}

	go http.ListenAndServe(":80", http.RedirectHandler("https://weaselfoss.dev", 200))

	fmt.Println("Started TLS server in Cert Manager mode.\nDBHost: ", os.Getenv("DATABASE_HOST"))
	err := server.ListenAndServeTLS("/root/resources/certs/fullchain1.pem", "/root/resources/certs/privkey1.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("Shutting down!")
}

func httpInterceptor(router http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(rw, req)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		fmt.Printf("Page load time: %v\n", elapsedTime)

	}
}
