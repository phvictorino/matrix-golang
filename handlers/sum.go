package handlers

import (
	"net/http"
	"test/services"
)

// SumHandler handle with invert request
func SumHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.SumService)
}
