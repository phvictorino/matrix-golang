package utils

// FlatMatrixToArray receive a matrix and returns a flatten array
func FlatMatrixToArray(matrix [][]int) []int {
	flattenArray := make([]int, 0)

	for _, row := range matrix {
		for _, column := range row {
			flattenArray = append(flattenArray, column)
		}
	}
	return flattenArray
}
