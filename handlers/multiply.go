package handlers

import (
	"net/http"
	"test/services"
)

// MultiplyHandler handle with echo request
func MultiplyHandler(response http.ResponseWriter, request *http.Request) {
	services.GenericService(response, request, services.MultiplyService)
}
