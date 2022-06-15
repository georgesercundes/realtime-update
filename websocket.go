package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWebsocket(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ws)
}

