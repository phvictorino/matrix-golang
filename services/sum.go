package services

import (
	"strconv"
	"test/utils"
)

// SumService sums whole values of a given matrix
func SumService(matrix [][]int, channel chan string) {
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 0
	for _, value := range flattenArray {
		sum += value
	}

	channel <- "Sum: " + strconv.Itoa(sum)
}
