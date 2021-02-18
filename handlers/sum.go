package handlers

import (
	"net/http"
	"test/services"
)

// SumHandler handle with echo request
func SumHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.SumService)
}
