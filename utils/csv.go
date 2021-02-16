package utils

import (
	"encoding/csv"
	"net/http"
	"strconv"
)

// CsvTransformer parse csv to matrix
func CsvTransformer(response http.ResponseWriter, request *http.Request) [][]int {
	reader := csv.NewReader(request.Body)
	result, error := reader.ReadAll()

	if error != nil {
		http.Error(response, "Error to convert CSV", http.StatusInternalServerError)
	}

	intMatrix := convertValuesFromStringToInt(result)

	return intMatrix
}

func convertValuesFromStringToInt(matrix [][]string) [][]int {
	newMatrix := make([][]int, len(matrix))
	InitializeRowsOfSlice(newMatrix)

	for indexRow, row := range matrix {
		for indexColumn, column := range row {
			result, _ := strconv.Atoi(column)
			newMatrix[indexRow][indexColumn] = result
		}
	}
	return newMatrix
}
