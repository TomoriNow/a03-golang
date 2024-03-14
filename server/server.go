package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	SERVER_HOST = ""
	SERVER_PORT = ""
	SERVER_TYPE = "tcp"
	BUFFER_SIZE = 1024
	GROUP_NAME  = ""
)

type HttpRequest struct {
	Method          string
	Uri             string
	Version         string
	Host            string
	Accept          string
	AcceptLanguange string
}

type HttpResponse struct {
	Version         string
	StatusCode      string
	ContentType     string
	ContentLanguage string
	Data            string
}

type Student struct {
	Nama string
	Npm  string
}

func main() {
	listenAddress, err := net.ResolveTCPAddr(SERVER_TYPE, net.JoinHostPort(SERVER_HOST, SERVER_PORT))
	if err != nil {
		log.Fatalln(err)
	}

	socket, err := net.ListenTCP(SERVER_TYPE, listenAddress)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("TCP Server Socket Program Example in Go that will connect to server\n")
	fmt.Printf("Press Ctrl+C or Cmd+C to stop the program\n")
	fmt.Printf("[%s] Listening on: %s\n", SERVER_TYPE, socket.Addr())

	defer socket.Close()

	for {
		connection, err := socket.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}

		go HandleConnection(connection)
	}
}

func HandleConnection(connection net.Conn) {
	fmt.Printf("[%s] Receive connection from %s\n", SERVER_TYPE, connection.RemoteAddr())
	fmt.Printf("[%s] [Client: %s] Creating receive buffer for connection of size %d\n", SERVER_TYPE, connection.RemoteAddr(), BUFFER_SIZE)
	receiveBuffer := make([]byte, BUFFER_SIZE)

	defer connection.Close()

	receiveLength, err := connection.Read(receiveBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[%s] [Client: %s] Received %d bytes of message\n", SERVER_TYPE, connection.RemoteAddr(), receiveLength)
	message := string(receiveBuffer[:receiveLength])

	fmt.Printf("[%s] [Client: %s] Message: %s\n", SERVER_TYPE, connection.RemoteAddr(), message)

	response, err := logic(message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[%s] [Client: %s] Sending Response: %s\n", SERVER_TYPE, connection.RemoteAddr(), response)
	_, err = connection.Write([]byte(response))
	if err != nil {
		log.Fatalln(err)
	}

}

func HandleRequest(req HttpRequest) HttpResponse {
	//This program handles the routing to each view handler.

}

func RequestDecoder(bytestream []byte) HttpRequest {
	//Put the decoding program for HTTP Request Packet here
	var req HttpRequest

	reqString := string(bytestream)

	lines := strings.Split(reqString, "\r\n")

	req.Method, req.Uri, req.Version = ExtractRequestLine(lines[0])

	hostLine := lines[1]
	parts := strings.Split(hostLine, ": ")
	req.Host = parts[1]

	acceptLine := lines[2]
	parts = strings.Split(acceptLine, ": ")
	req.Accept = parts[1]

	acceptLanguageLine := lines[3]
	parts = strings.Split(acceptLanguageLine, ": ")
	req.AcceptLanguange = parts[1]

	return req

}

func ExtractRequestLine(requestLine string) (string, string, string) {
	parts := strings.Split(requestLine, " ")
	return parts[0], parts[1], parts[2]
}

func ResponseEncoder(res HttpResponse) []byte {
	//Put the encoding program for HTTP Response Struct here
	var result string

	return []byte(result)

}
func logic(input string) (string, error) {
	return strings.ToUpper(input), nil
}
