package services

import (
	"fmt"
	"net/http"
	"test/utils"
)

// GenericService receives the request and the service function.
// It calls the CsvTransform before all services
func GenericService(response http.ResponseWriter, request *http.Request, serviceToExecute func(matrix [][]int, channel chan string)) {
	// Prepare channel and matrix
	channel := make(chan string)
	matrix := utils.CsvTransformer(response, request)

	// Call service received and wait for the channel response
	go serviceToExecute(matrix, channel)
	serviceResponse := <-channel

	// Print on console and return the service response
	fmt.Println(serviceResponse)
	response.Write([]byte(serviceResponse))
}
