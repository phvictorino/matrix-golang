package utils

// InitializeRowsOfSlice receive a slice and initialize rows with empty arrays
func InitializeRowsOfSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, len(slice))
	}
}
