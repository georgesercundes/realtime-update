package main

import "net/http"

func setupRoutes() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(hub, w, r)
	})
}
