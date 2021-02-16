package handlers

import (
	"net/http"
	"strconv"
	"test/utils"
)

// SumHandler sums whole values of a given matrix
func SumHandler(response http.ResponseWriter, request *http.Request) {
	matrix := utils.CsvTransformer(response, request)
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 0
	for _, value := range flattenArray {
		sum += value
	}

	sumConvertedToString := strconv.Itoa(sum)
	response.Write([]byte(sumConvertedToString))
}
