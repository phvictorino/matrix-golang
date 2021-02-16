package handlers

import (
	"net/http"
	"test/matrixstring"
	"test/utils"
)

// InvertHandler inverts the matrix
func InvertHandler(response http.ResponseWriter, request *http.Request) {
	matrix := utils.CsvTransformer(response, request)
	newArray := make([][]int, len(matrix))
	utils.InitializeRowsOfSlice(newArray)

	for indexRow, row := range matrix {
		for indexColumn, column := range row {
			newArray[indexColumn][indexRow] = column
		}
	}

	formattedSlice := matrixstring.TransformMatrixToString(newArray)
	response.Write([]byte(formattedSlice))

}
