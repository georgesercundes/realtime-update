package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Data
}

func newHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Data),
	}
}

func (h *Hub) run() {
	for {
		select {
		case message := <-h.broadcast:
			for client := range h.clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("error occurred while writing the JSON -  %v", err)
				}
			}
		}
	}
}
