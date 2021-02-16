package main

import (
	"encoding/csv"
	"net/http"
)

// CsvHandler parse csv to matrix
func CsvHandler(response http.ResponseWriter, request *http.Request) [][]string {
	reader := csv.NewReader(request.Body)
	result, error := reader.ReadAll()

	if error != nil {
		http.Error(response, "Error to convert CSV", http.StatusInternalServerError)
	}

	return result
}
