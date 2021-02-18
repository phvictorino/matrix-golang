package handlers

import (
	"net/http"
	"test/services"
)

// EchoHandler handle with echo request
func EchoHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.EchoService)
}
