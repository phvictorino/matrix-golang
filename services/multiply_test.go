package services

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	channel := make(chan string)
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "Multiply: 3628800"
	go MultiplyService(input, channel)
	result := <-channel
	if expectedResult != result {
		t.Errorf("Multiply result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
