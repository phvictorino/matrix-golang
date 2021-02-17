package services

import (
	"test/matrixstring"
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

	formattedSlice := matrixstring.TransformMatrixToString(newArray)
	return formattedSlice

}
