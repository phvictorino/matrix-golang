package services

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "1,2,3,4,5,6,7,8,9"
	result := FlattenService(input)
	if expectedResult != result {
		t.Errorf("Flatten result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
