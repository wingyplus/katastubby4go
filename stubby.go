package stubby4go

import (
	"fmt"
	"net/http"
)

type Request struct {
	Method string
	Url    string
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

func CreateHandler(request Request, response Response) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(responseWriter, response.Body)
			responseWriter.Header().Set("content-type", response.Headers["content-type"])
		},
	)
}
