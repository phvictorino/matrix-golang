package main

import (
	"fmt"
	"net/http"
)

// InvertHandler inverts the matrix
func InvertHandler(response http.ResponseWriter, request *http.Request) {
	matrix := CsvHandler(response, request)

	for index, element := range matrix {
		fmt.Println(index, element)
	}
}
