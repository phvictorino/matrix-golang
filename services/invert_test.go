package services

import (
	"testing"
)

func TestInvert(t *testing.T) {
	channel := make(chan string)
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "Invert:\n1,4,7\n2,5,8\n3,6,9\n"
	go InvertService(input, channel)
	result := <-channel
	if expectedResult != result {
		t.Errorf("Invert result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
