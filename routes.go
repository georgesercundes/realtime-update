package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("/websocket", serveWebsocket)
}