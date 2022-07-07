package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	wsClientsHub := newWSHub()
	go wsClientsHub.run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(wsClientsHub, w, r)
	})

	fmt.Println("listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
