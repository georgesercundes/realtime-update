package main

import (
	"fmt"
	"log"
	"net/http"
)

// SSEChannel model

func main() {
	SSEClientsHub := newSSEHub()

	go SSEClientsHub.run()

	http.HandleFunc("/sse", func(rw http.ResponseWriter, r *http.Request) {
		write(SSEClientsHub, rw, r)
	})

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		read(SSEClientsHub, rw, r)
	})

	fmt.Println("listening on http://127.0.0.1:5000")
	log.Fatal(http.ListenAndServeTLS(":5000", "../../cert/server.crt", "../../cert/server.key", nil))
}
