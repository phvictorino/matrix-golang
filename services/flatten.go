package services

import (
	"test/utils"
)

// FlattenService flat given array and return elements separated by comma
func FlattenService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)

	// Used this just to respond the http request
	return utils.TransformArrayToContinuousString(flattenArray)
}
