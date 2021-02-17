package handlers

import (
	"net/http"
	"test/services"
)

// InvertHandler handle with invert request
func InvertHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.InvertService)
}
