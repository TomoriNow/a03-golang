package main

import (
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
	//The Program logic should go here.

}

func HandleConnection(connection net.Conn) {
	//This progrom handles the incoming request from client

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

func ExtractRequestLine (requestLine string) (string, string, string) {
	parts := strings.Split(requestLine, " ")
	return parts[0], parts[1], parts[2]
}

func ResponseEncoder(res HttpResponse) []byte {
	//Put the encoding program for HTTP Response Struct here
	var result string

	return []byte(result)

}
