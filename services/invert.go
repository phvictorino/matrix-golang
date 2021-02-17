package services

import (
	"test/utils"
)

// InvertService inverts the matrix
func InvertService(matrix [][]int) string {

	newArray := make([][]int, len(matrix))
	utils.InitializeRowsOfSlice(newArray)

	for indexRow, row := range matrix {
		for indexColumn, column := range row {
			newArray[indexColumn][indexRow] = column
		}
	}

	// Used this just to respond the http request
	formattedSlice := utils.TransformMatrixToString(newArray)
	return formattedSlice

}
