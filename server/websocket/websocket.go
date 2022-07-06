package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime_update/server/data"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	wsChannel := newWSChannel()
	go wsChannel.run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(wsChannel, w, r)
	})

	fmt.Println("listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveWebsocket(channel *WSChannel, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	messageChannel := make(chan data.Data)

	client := &Client{channel: channel, conn: ws, send: messageChannel}
	client.channel.register <- client

	go client.read()
	go client.write()
}
