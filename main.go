package main

import (
	"log"
	"net/http"
	"test/handlers"
)

func main() {
	http.HandleFunc("/invert", handlers.InvertHandler)
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/flatten", handlers.FlattenHandler)
	http.HandleFunc("/sum", handlers.SumHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
