package main

import (
	"log"
	"strings"
)

type Request struct {
	Method  string
	Version string
	Path    string
}

func NewRequest(requestString string) *Request {

	requestParts := strings.Split(requestString, "\n")

	request := Request{}

	requestInfo := strings.Split(requestParts[0], " ")

	if len(requestInfo) < 2 {
		log.Fatalln("Something went wrong")
	}

	request.Method = requestInfo[0]

	request.Path = parsePath(requestInfo[1])

	request.Version = "HTTP/1.1"

	return &request

}

func parsePath(path string) string {

	if path == "/" {
		return "/index.html"
	}

	return path

}
