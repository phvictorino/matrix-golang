package services

import (
	"strconv"
	"test/utils"
)

// MultiplyService multiply whole values of a given matrix
func MultiplyService(matrix [][]int) string {
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 1
	for _, value := range flattenArray {
		sum += value * sum
	}

	return strconv.Itoa(sum)

}
