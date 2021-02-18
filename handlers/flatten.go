package handlers

import (
	"net/http"
	"test/services"
)

// FlattenHandler handle with echo request
func FlattenHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.FlattenService)
}
