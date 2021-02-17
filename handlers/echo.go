package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"test/services"
	"test/utils"
	"time"
)

// EchoHandler handle with echo request
func EchoHandler(response http.ResponseWriter, request *http.Request) {
	channel := make(chan string)

	matrix := utils.CsvTransformer(response, request)

	go services.EchoService(matrix, channel)
	go services.InvertService(matrix, channel)
	go services.FlattenService(matrix, channel)
	go services.SumService(matrix, channel)
	go services.MultiplyService(matrix, channel)

	time.AfterFunc(time.Second*2, func() { close(channel) })

	var sb strings.Builder
	for channelResponse := range channel {
		fmt.Println(channelResponse)
		sb.WriteString(channelResponse + "\n")
	}
	response.Write([]byte(sb.String()))
}
