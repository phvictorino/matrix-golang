package services

import (
	"test/matrixstring"
	"test/utils"
)

// FlattenService flat given array
func FlattenService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)
	return matrixstring.TransformArrayToContinuousString(flattenArray)
}
