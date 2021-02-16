package main

import (
	"net/http"
	"strings"
)

// EchoHandler echoes the matrix
func EchoHandler(response http.ResponseWriter, request *http.Request) {
	matrix := CsvHandler(response, request)
	var result strings.Builder
	for index, element := range matrix {
		result.WriteString(strings.Join(element, ","))
		if index < len(element)-1 {
			result.WriteString(",")
		}
	}
	response.Write([]byte(result.String()))
}
