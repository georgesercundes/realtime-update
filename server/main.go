package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupRoutes()
	fmt.Println("listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
