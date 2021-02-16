package handlers

import (
	"net/http"
	"strconv"
	"test/utils"
)

// MultiplyHandler multiply whole values of a given matrix
func MultiplyHandler(response http.ResponseWriter, request *http.Request) {
	matrix := utils.CsvTransformer(response, request)
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 1
	for _, value := range flattenArray {
		sum += value * sum
	}

	sumConvertedToString := strconv.Itoa(sum)
	response.Write([]byte(sumConvertedToString))
}
