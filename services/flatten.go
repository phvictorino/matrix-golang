package services

import (
	"test/utils"
)

// FlattenService flat given array
func FlattenService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)
	return utils.TransformArrayToContinuousString(flattenArray)
}
