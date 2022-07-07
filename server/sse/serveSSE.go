package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"realtime_update/server/data"
)

func read(hub *SSEClientsHub, rw http.ResponseWriter, r *http.Request) {
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
	hub.broadcast <- message
}

func write(hub *SSEClientsHub, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Connection doesn't support streaming", http.StatusBadRequest)
		return
	}

	messageChannel := make(chan data.Data)
	client := &Client{hub: hub, send: messageChannel}
	client.hub.register <- client

	notify := rw.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notify:
			client.hub.unregister <- client
			return
		case data := <-client.send:
			fmt.Fprintf(rw, "data: %v \n\n", data.Message)
			flusher.Flush()
		}
	}
}
