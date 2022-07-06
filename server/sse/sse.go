package main

import (
	"fmt"
	"log"
	"net/http"
)

// SSEChannel model

func main() {
	sseChannel := newSSEChannel()

	go sseChannel.run()

	http.HandleFunc("/sse", sseChannel.write)

	http.HandleFunc("/log", sseChannel.read)

	fmt.Println("listening on http://127.0.0.1:5000")
	log.Fatal(http.ListenAndServeTLS(":5000", "../../cert/server.crt", "../../cert/server.key", nil))
}
