package main

import (
	"fmt"
	"net"
)

type HttpRequest struct {
	Method         string
	Uri            string
	Version        string
	Host           string
	Accept         string
	AcceptLanguage string
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

const (
	SERVER_TYPE = "tcp"
	BUFFER_SIZE = 1024
)

func main() {
	//The Program logic should go here.

}

func Fetch(req HttpRequest, connection net.Conn) (HttpResponse, []Student, HttpRequest) {
	//This program handles the request-making to the server
	var res HttpResponse
	var Student []Student

	return res, Student, req

}

func ResponseDecoder(bytestream []byte) HttpResponse {
	var res HttpResponse

	return res

}

func RequestEncoder(req HttpRequest) []byte {
	var result string

	// Construct the request line
	requestLine := fmt.Sprintf("%s %s %s", req.Method, req.Uri, req.Version)

	// Construct other header fields
	headers := fmt.Sprintf("Host: %s\r\nAccept: %s\r\nAccept-Language: %s\r\n", req.Host, req.Accept, req.AcceptLanguage)

	// Combine the request line, headers, and an empty line to separate header and data
	result = fmt.Sprintf("%s\r\n%s\r\n", requestLine, headers)

	return []byte(result)

}
