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
		http.Error(response, "Error to open File", http.StatusInternalServerError)
	}
	defer file.Close()

	result, error := csv.NewReader(file).ReadAll()

	if error != nil {
		http.Error(response, "Error to convert file to CSV", http.StatusInternalServerError)
		return nil
	}

	validatedMatrix, error := validateMatrixAndCovertElementsToInt(result)

	if error != nil {
		http.Error(response, error.Error(), http.StatusBadRequest)
		return nil
	}

	return validatedMatrix
}

func validateMatrixAndCovertElementsToInt(matrix [][]string) ([][]int, error) {
	squareError := validateSquareMatrix(matrix)

	if squareError != nil {
		return nil, squareError
	}

	intMatrix, convertError := convertValuesFromStringToInt(matrix)

	if convertError != nil {
		return nil, convertError
	}

	return intMatrix, nil
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

func validateSquareMatrix(matrix [][]string) error {
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
