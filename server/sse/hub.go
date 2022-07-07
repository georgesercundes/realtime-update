package main

import (
	"realtime_update/server/data"
)

type Client struct {
	hub  *SSEClientsHub
	send chan data.Data
}

type SSEClientsHub struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan data.Data
}

func newSSEHub() *SSEClientsHub {
	return &SSEClientsHub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan data.Data),
	}
}

func (hub *SSEClientsHub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
		case data := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- data:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}
