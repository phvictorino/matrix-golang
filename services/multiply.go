package services

import (
	"strconv"
	"test/utils"
)

// MultiplyService multiply whole values of a given matrix
func MultiplyService(matrix [][]int, channel chan string) {
	flattenArray := utils.FlatMatrixToArray(matrix)

	sum := 1
	for _, value := range flattenArray {
		sum += value * sum
	}

	channel <- "Multiply: " + strconv.Itoa(sum)
}
