package services

import (
	"test/matrixstring"
)

// EchoService echoes the matrix
func EchoService(matrix [][]int) string {
	return matrixstring.TransformMatrixToString(matrix)
}
