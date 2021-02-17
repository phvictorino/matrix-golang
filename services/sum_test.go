package services

import (
	"testing"
)

func TestSum(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "45"
	result := SumService(input)
	if expectedResult != result {
		t.Errorf("Multiply result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
