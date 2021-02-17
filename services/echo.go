package services

import (
	"test/utils"
)

// EchoService just echoes the matrix with commas between numbers
func EchoService(matrix [][]int) string {
	return utils.TransformMatrixToString(matrix)
}
