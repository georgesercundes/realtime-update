package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	defer func() {
		delete(hub.clients, ws)
		ws.Close()
		log.Println("Closed!")
	}()

	hub.clients[ws] = true

	log.Println("Connected")

	read(hub, ws)
}

func read(hub *Hub, client *websocket.Conn) {
	for {
		var message Data

		err := client.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading JSON - %v", err)
			delete(hub.clients, client)
			return
		}

		hub.broadcast <- message
	}
}
