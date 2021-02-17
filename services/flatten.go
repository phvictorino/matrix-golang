package services

import (
	"test/utils"
)

// FlattenService flat given array and return elements separated by comma
func FlattenService(matrix [][]int, channel chan string) {
	flattenArray := utils.FlatMatrixToArray(matrix)
	channel <- "Flatten: " + utils.TransformArrayToContinuousString(flattenArray)
}
