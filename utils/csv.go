package utils

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
)

// CsvTransformer parse csv to matrix
func CsvTransformer(response http.ResponseWriter, request *http.Request) [][]int {

	file, _, err := request.FormFile("file")
	if err != nil {
		http.Error(response, "Error to convert CSV", http.StatusInternalServerError)
	}
	defer file.Close()

	result, error := csv.NewReader(file).ReadAll()

	if error != nil {
		http.Error(response, "Error to convert CSV", http.StatusInternalServerError)
	}

	intMatrix, convertError := convertValuesFromStringToInt(result)

	if convertError != nil {
		http.Error(response, convertError.Error(), http.StatusBadRequest)
	}

	squareError := validateSquareMatrix(intMatrix)

	if squareError != nil {
		http.Error(response, squareError.Error(), http.StatusInternalServerError)
	}

	return intMatrix
}

func convertValuesFromStringToInt(matrix [][]string) ([][]int, error) {
	newMatrix := make([][]int, len(matrix))
	InitializeRowsOfSlice(newMatrix)

	for indexRow, row := range matrix {
		for indexColumn, column := range row {
			result, error := strconv.Atoi(column)
			if (error) != nil {
				return nil, error
			}
			newMatrix[indexRow][indexColumn] = result
		}
	}
	return newMatrix, nil
}

func validateSquareMatrix(matrix [][]int) error {
	var lengthRow, lengthColumn int
	lengthRow = len(matrix)
	for _, row := range matrix {
		lengthColumn = len(row)
	}

	if lengthColumn != lengthRow {
		return errors.New("Matrix must be square: same length of rows and columns")
	}
	return nil
}
