package main

import (
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

func serveWebsocket(hub *WSClientsHub, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	messageChannel := make(chan data.Data)

	client := &Client{hub: hub, conn: ws, send: messageChannel}
	client.hub.register <- client

	go read(client)
	go write(client)
}

func read(client *Client) {
	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()

	for {
		var message data.Data
		err := client.conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading JSON - %v", err)
			break
		}

		client.hub.broadcast <- message
	}
}

func write(client *Client) {
	defer func() {
		client.conn.Close()
	}()

	for {
		select {
		case message := <-client.send:
			err := client.conn.WriteJSON(message)
			if err != nil {
				log.Printf("error occurred while writing JSON - %v", err)
				return
			}
		}
	}
}
