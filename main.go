package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Trey2k/OpenStreaming/app/home"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	http.HandleFunc("/", httpInterceptor(router))
	router.HandleFunc("/", home.GetHomePage).Methods("GET")
	router.HandleFunc("/twitch", home.TwitchOAuthEndpoint()).Methods("GET")

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
