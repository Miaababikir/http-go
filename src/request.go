package main

import (
	"strings"
)

type Request struct {
	Method  string
	Version string
	Path    string
	Headers map[string]string
}

func NewRequest(requestString string) *Request {

	requestParts := strings.SplitAfter(requestString, "\r\n")

	request := parseRequest(requestParts[0])

	request.Headers = parseHeaders(requestParts[1:])

	return request
}

func parseHeaders(headerString []string) map[string]string {

	headers := make(map[string]string)

	for _, value := range headerString {

		if value == "\r\n" {
			break
		}

		parts := strings.Split(value, ":")
		headers[parts[0]] = parts[1]

	}

	return headers
}

func parseRequest(statusLine string) *Request {
	requestDetails := strings.Split(statusLine, " ")

	request := &Request{
		Method:  requestDetails[0],
		Path:    parsePath(requestDetails[1]),
		Version: "HTTP/1.1",
	}

	return request
}

func parsePath(path string) string {

	if path == "/" {
		return "/index.html"
	}

	return path

}
