package services

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	channel := make(chan string)
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "Flatten: 1,2,3,4,5,6,7,8,9\n"
	go FlattenService(input, channel)
	serviceResult := <-channel

	if expectedResult != serviceResult {
		t.Errorf("Flatten result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, serviceResult)
	}
}
