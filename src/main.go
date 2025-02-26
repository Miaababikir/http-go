package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Print("Http server using TCP from scratch\n")

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal("Not able to start a server")
	}

	defer listener.Close()

	fmt.Printf("Listening at Port %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Failed to accept the connection: %s", err)
		}

		reader := bufio.NewReader(conn)

		var requestBuf bytes.Buffer

		for {
			line, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal("Something went wrong")
			}

			requestBuf.WriteString(line)

			if strings.HasSuffix(requestBuf.String(), "\r\n\r\n") {
				break
			}
		}

		request := NewRequest(requestBuf.String())

		response := NewResponse(request)

		log.Printf("Request for %s, respond with %d status\n\n", request.Path, response.Code)

		conn.Write(response.ToBytes())

		conn.Close()

	}
}
