package handlers

import (
	"net/http"
	"test/matrixstring"
	"test/utils"
)

// FlattenHandler flat given array
func FlattenHandler(response http.ResponseWriter, request *http.Request) {
	matrix := CsvTransformer(response, request)
	flattenArray := utils.FlatMatrixToArray(matrix)
	resultString := matrixstring.TransformArrayToContinuousString(flattenArray)
	response.Write([]byte(resultString))
}
