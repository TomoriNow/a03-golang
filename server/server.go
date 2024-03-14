package main

import (
	"encoding/json"
	"encoding/xml"
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
	//The Program logic should go here.

}

func HandleConnection(connection net.Conn) {
	//This progrom handles the incoming request from client
	receiveBuffer := make([]byte, BUFFER_SIZE)

	defer connection.Close()

	receiveLength, err := connection.Read(receiveBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	rawRequest := receiveBuffer[:receiveLength]

	request := RequestDecoder(rawRequest)

	response := HandleRequest(request)


}

func HandleRequest(req HttpRequest) HttpResponse {
	var response HttpResponse
	//This program handles the routing to each view handler.
	response.StatusCode = "404"
	if req.Uri == "/" || req.Uri == "/?name="+GROUP_NAME {
		response.Data = "<html><body><h1>Halo, kami dari Klepon</h1></body>>/html>"
		response.StatusCode = "200"
		response.ContentType = "text/html"
	}
	if req.Uri == "/data" {
		students := []Student{
			{Nama: "Sean", Npm: "2206822963"},
			{Nama: "Gusti", Npm: "2206821241"},
			{Nama: "Galih", Npm: "2206046696"},
		}
		if strings.Contains(req.Accept, "xml") {
			xmlData, err := xml.Marshal(students)
			if err != nil {
				fmt.Println("Error:", err)
			}
			response.StatusCode = "200"
			response.ContentType = "application/xml"
			response.Data = string(xmlData)
		} else if strings.Contains(req.Accept, "json") || strings.Contains(req.Accept, "q=") {
			jsonData, err := json.Marshal(students)
			if err != nil {
				fmt.Println("Error:", err)
			}
			response.StatusCode = "200"
			response.ContentType = "application/json"
			response.Data = string(jsonData)
		}
	}
	if req.Uri == "/greeting" {
		if strings.Contains(req.AcceptLanguange, "id") {
			response.Data = "<html><body><h1>Halo, kami dari Klepon</h1></body>>/html>"
			response.ContentType = "text/html"
			response.StatusCode = "200"
		} else if strings.Contains(req.AcceptLanguange, "en") || strings.Contains(req.Accept, "q=") {
			response.Data = "<html><body><h1>Hello, we are from Klepon</h1></body>>/html>"
			response.ContentType = "text/html"
			response.StatusCode = "200"
		}
	}

	if response.StatusCode != "404" {
		response.ContentLanguage = req.AcceptLanguange
		response.Version = req.Version
	}

	return response
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
