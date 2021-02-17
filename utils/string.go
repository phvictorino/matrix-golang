package utils

import (
	"strconv"
	"strings"
)

// TransformArrayToContinuousString receive a array and returns an one line string | e.g: 1, 2, 3, 4, 5, 6
func TransformArrayToContinuousString(array []int) string {
	var result strings.Builder
	for index, element := range array {
		result.WriteString(strconv.Itoa(element))
		if index < len(array)-1 {
			result.WriteString(",")
		}
	}
	return result.String()
}

// TransformMatrixToString receive a matrix and return formatted string with break lines
func TransformMatrixToString(slice [][]int) string {
	var result strings.Builder
	for _, element := range slice {
		for index2, element2 := range element {
			result.WriteString(strconv.Itoa(element2))
			if index2 < len(slice)-1 {
				result.WriteString(",")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}
