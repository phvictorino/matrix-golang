package services

import (
	"testing"
)

func TestEcho(t *testing.T) {

	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedResult := "1,2,3\n4,5,6\n7,8,9\n"
	EchoService(input, nil)

	if expectedResult != "result" {
		t.Errorf("Echo result doesn't match")
		t.Errorf("Expected: %s | Result: %s", expectedResult, "result")
	}
}
