package services

import (
	"test/utils"
)

// EchoService echoes the matrix with commas between numbers
func EchoService(matrix [][]int, channel chan string) {
	channel <- "Echo: " + utils.TransformMatrixToString(matrix)
}
