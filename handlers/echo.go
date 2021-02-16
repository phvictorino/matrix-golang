package handlers

import (
	"net/http"
	"test/matrixstring"
	"test/utils"
)

// EchoHandler echoes the matrix
func EchoHandler(response http.ResponseWriter, request *http.Request) {
	matrix := utils.CsvTransformer(response, request)
	stringResult := matrixstring.TransformMatrixToString(matrix)
	response.Write([]byte(stringResult))
}
