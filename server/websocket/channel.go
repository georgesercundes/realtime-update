package main

import (
	"log"
	"realtime_update/server/data"

	"github.com/gorilla/websocket"
)

type Client struct {
	channel *WSChannel
	conn    *websocket.Conn
	send    chan data.Data
}

type WSChannel struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan data.Data
}

func newWSChannel() *WSChannel {
	return &WSChannel{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan data.Data),
	}
}

func (channel *WSChannel) run() {
	for {
		select {
		case client := <-channel.register:
			channel.clients[client] = true
		case client := <-channel.unregister:
			delete(channel.clients, client)
		case message := <-channel.broadcast:
			for client := range channel.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(channel.clients, client)
				}
			}
		}
	}
}

func (c *Client) read() {
	defer func() {
		c.channel.unregister <- c
		c.conn.Close()
	}()

	for {
		var message data.Data
		err := c.conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading JSON - %v", err)
			break
		}

		c.channel.broadcast <- message
	}
}

func (c *Client) write() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message := <-c.send:
			err := c.conn.WriteJSON(message)
			if err != nil {
				log.Printf("error occurred while writing JSON - %v", err)
				return
			}
		}
	}
}
