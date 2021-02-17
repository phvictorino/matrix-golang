package services

import (
	"testing"
)

func TestSum(t *testing.T) {
	channel := make(chan string)
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "Sum: 45"
	go SumService(input, channel)
	result := <-channel
	if expectedResult != result {
		t.Errorf("Sum result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
