package main

import (
	"log"
	"net/http"
	"test/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/invert", handlers.InvertHandler)
	http.HandleFunc("/flatten", handlers.FlattenHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)
	http.HandleFunc("/sum", handlers.SumHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
