package services

import (
	"test/utils"
)

// EchoService echoes the matrix
func EchoService(matrix [][]int) string {
	return utils.TransformMatrixToString(matrix)
}
