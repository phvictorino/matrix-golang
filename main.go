package main

import (
	"log"
	"net/http"
	"test/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.EchoHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
