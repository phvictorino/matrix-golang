package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/echo", EchoHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
