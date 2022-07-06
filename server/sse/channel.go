package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"realtime_update/server/data"
)

type Client struct {
	channel *SSEChannel
	send    chan data.Data
}

type SSEChannel struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan data.Data
}

func newSSEChannel() *SSEChannel {
	return &SSEChannel{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan data.Data),
	}
}

func (channel *SSEChannel) run() {
	for {
		select {
		case client := <-channel.register:
			channel.clients[client] = true
		case client := <-channel.unregister:
			delete(channel.clients, client)
			close(client.send)
		case data := <-channel.broadcast:
			for client := range channel.clients {
				client.send <- data
			}
		}
	}
}

func (channel *SSEChannel) read(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read request body: %s", err)
		return
	}

	var message data.Data

	err = json.Unmarshal(body, &message)
	if err != nil {
		log.Printf("Could not unmarshall JSON to data structure: %s", err)
		return
	}
	channel.broadcast <- message
}

func (channel *SSEChannel) write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Connection doesn't support streaming", http.StatusBadRequest)
		return
	}

	messageChannel := make(chan data.Data)
	client := &Client{channel: channel, send: messageChannel}
	channel.clients[client] = true

	notify := w.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notify:
			client.channel.unregister <- client
			return
		case data := <-messageChannel:
			fmt.Fprintf(w, "data: %v \n\n", data.Message)
			flusher.Flush()
		}
	}
}
