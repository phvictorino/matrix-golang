package services

import (
	"strconv"
	"test/utils"
)

// SumService sums whole values of a given matrix
func SumService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 0
	for _, value := range flattenArray {
		sum += value
	}

	return strconv.Itoa(sum)
}
