package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trey2k/OpenStreaming/webApp/app/api"
	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/Trey2k/OpenStreaming/webApp/app/dashboard"
	"github.com/Trey2k/OpenStreaming/webApp/app/overlay"
	"github.com/Trey2k/OpenStreaming/webApp/app/twitch/eventSub"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.HandleFunc("/", httpInterceptor(router))
	router.Handle("/", http.RedirectHandler("/dashboard", http.StatusFound)).Methods("GET")

	// Dashboard
	router.HandleFunc("/dashboard", dashboard.AuthenticatedMW(dashboard.GetHomePage)).Methods("GET")
	router.HandleFunc("/login", dashboard.GetLoginPage).Methods("GET")
	router.HandleFunc("/twitch", dashboard.TwitchOAuthEndpoint()).Methods("GET")

	// Overlay
	router.HandleFunc("/overlay/{id}", overlay.OverlayHandler).Methods("GET")
	// Overlay
	router.HandleFunc("/overlay/{id}/editor", dashboard.AuthenticatedMW(overlay.OverlayEditorHandler)).Methods("GET")

	// Api endpoints
	router.HandleFunc("/api/getEvents", api.GetEventHandler).Methods("GET")
	router.HandleFunc("/api/toggleBot", api.ToggleBotHandler).Methods("GET")

	// Overlay Api endpoints
	router.HandleFunc("/api/overlay/websocket", api.OverlayWSHandler)

	// Favicon handler
	router.HandleFunc("/favicon.ico", faviconHandler)

	// Overlay Wasm handler
	router.HandleFunc("/overlay.wasm", overlayWasmHandler)

	// Static file server
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("/root/resources/static")))
	http.Handle("/static/", fileServer)

	// Callback
	router.HandleFunc("/eventsub", eventSub.EventSubHandler()).Methods("POST")

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PortSSL")),
	}

	go http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("Port")), http.RedirectHandler(fmt.Sprintf("%s/dashboard", os.Getenv("URL")), http.StatusFound))

	common.Loggers.Info.Printf("Started TLS server\n")
	err := server.ListenAndServeTLS(fmt.Sprintf("/root/%s", os.Getenv("FullChain")), fmt.Sprintf("/root/%s", os.Getenv("PrivateKey")))
	if err != nil {
		panic(err)
	}
}

func httpInterceptor(router http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(rw, req)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/root/resources/OpenStreaming.ico")
}

func overlayWasmHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/root/overlay.wasm")
}
