package main

import (
	"realtime_update/server/data"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *WSClientsHub
	conn *websocket.Conn
	send chan data.Data
}

type WSClientsHub struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan data.Data
}

func newWSHub() *WSClientsHub {
	return &WSClientsHub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan data.Data),
	}
}

func (hub *WSClientsHub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}
