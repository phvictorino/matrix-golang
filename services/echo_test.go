package services

import (
	"testing"
)

func TestEcho(t *testing.T) {
	channel := make(chan string)
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "Echo:\n1,2,3\n4,5,6\n7,8,9\n"
	go EchoService(input, channel)
	result := <-channel

	if expectedResult != result {
		t.Errorf("Echo result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, result)
	}
}
