package services

import (
	"net/http"
	"test/utils"
)

// GenericService receive the http request with csv and returns the converted matrix
func GenericService(response http.ResponseWriter, request *http.Request, serviceFunction func([][]int) string) {
	matrix := utils.CsvTransformer(response, request)
	serviceResponse := serviceFunction(matrix)
	response.Write([]byte(serviceResponse))
}
