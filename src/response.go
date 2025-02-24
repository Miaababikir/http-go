package main

import (
	"fmt"
	"os"
	"strings"
)

type Response struct {
	Version string
	Code    int
	Message string
	Body    string
}

func NewResponse(request *Request) *Response {

	response := Response{
		Version: request.Version,
		Code:    200,
		Message: "Success",
	}

	if request.Method == "GET" {
		fileName := request.Path
		data, err := os.ReadFile("./static" + fileName)

		if err != nil {
			return &Response{
				Version: request.Version,
				Code:    400,
				Message: "Not found",
			}
		}

		response.Body = string(data)
	}

	return &response
}

func NewNotFound() *Response {
	return &Response{
		Version: "HTTP/1.1",
		Code:    400,
		Message: "Not found",
	}
}

func (response *Response) ToBytes() []byte {

	responseInfo := fmt.Sprintf("%s %d %s\r\n\r\n", response.Version, response.Code, response.Message)

	body := response.Body

	var builder strings.Builder

	builder.WriteString(responseInfo)

	if response.Body != "" {
		builder.WriteString(body)
	}

	return []byte(builder.String())
}
