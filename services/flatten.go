package services

import (
	"test/utils"
)

// FlattenService flat given array and return elements separated by comma
func FlattenService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)
	return utils.TransformArrayToContinuousString(flattenArray)
}
