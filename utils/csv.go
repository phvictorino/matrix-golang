package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
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

	intMatrix, convertError := convertValuesFromStringToInt(result)

	if convertError != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(convertError.Error()))
	}

	return intMatrix
}

func convertValuesFromStringToInt(matrix [][]string) ([][]int, error) {
	squareError := validateSquareMatrix(matrix)

	if squareError != nil {
		return nil, squareError
	}

	newMatrix := make([][]int, len(matrix))
	InitializeRowsOfSlice(newMatrix)

	for indexRow, row := range matrix {
		for indexColumn, column := range row {
			result, error := strconv.Atoi(column)
			if (error) != nil {
				return nil, fmt.Errorf("Cannot cast element of the matrix to int: %v", error)
			}
			newMatrix[indexRow][indexColumn] = result
		}
	}
	return newMatrix, nil
}

func validateSquareMatrix(matrix [][]string) error {
	var lengthRow, lengthColumn int
	for _, row := range matrix {
		lengthRow = len(row)
		for _, column := range row {
			lengthColumn = len(column)
		}
	}

	if lengthColumn != lengthRow {
		return errors.New("Matrix must be square: same length of rows and columns")
	}
	return nil
}
