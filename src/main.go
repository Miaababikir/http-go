package main

import (
	"fmt"
	"log"
	"net"
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

		var requestBuffer = make([]byte, 1024)

		conn.Read(requestBuffer)

		request := NewRequest(string(requestBuffer))

		response := NewResponse(request)

		log.Printf("Request for %s, respond with %d status\n\n", request.Path, response.Code)

		conn.Write(response.ToBytes())

		conn.Close()

	}
}
