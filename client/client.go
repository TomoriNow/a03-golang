package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
type Students struct {
	XMLName xml.Name
	Student []Student
}

const (
	SERVER_TYPE = "tcp"
	BUFFER_SIZE = 1024
)

func main() {

	fmt.Printf("[%s] Creating receive buffer of size %d\n", SERVER_TYPE, BUFFER_SIZE)
	fmt.Print("input the url: ")
	url, err := bufio.NewReader(os.Stdin).ReadString('\n')
	url = strings.TrimRight(url, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("input the data type: ")
	mimeType, err := bufio.NewReader(os.Stdin).ReadString('\n')
	mimeType = strings.TrimRight(mimeType, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Input the language: ")
	language, err := bufio.NewReader(os.Stdin).ReadString('\n')
	language = strings.TrimRight(language, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	host, port := getHostAndPortFromUrl(url)
	uri := getURI(url)

	request := HttpRequest{
		Method:         "GET",
		Uri:            uri,
		Version:        "HTTP/1.1",
		Host:           host,
		Accept:         mimeType,
		AcceptLanguage: language,
	}

	remoteTcpAddress, err := net.ResolveTCPAddr(SERVER_TYPE, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalln(err)
	}
	socket, err := net.DialTCP(SERVER_TYPE, nil, remoteTcpAddress)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("HTTP/1.1 Client-Server Program in Go\n")
	fmt.Printf("[%s] Dialling from %s to %s\n", SERVER_TYPE, socket.LocalAddr(), socket.RemoteAddr())

	defer socket.Close()

	response, _, _ := Fetch(request, socket)
	fmt.Println("Status: ", response.StatusCode)
	fmt.Print("Body:", response.Data)
}

func getHostAndPortFromUrl(url string) (string, string) {
	parts := strings.Split(url, "://")

	if len(parts) > 1 {

		hostAndPort := strings.Split(parts[1], "/")[0]
		hostPortParts := strings.Split(hostAndPort, ":")

		if len(hostPortParts) == 2 {
			return hostPortParts[0], hostPortParts[1]
		} else {
			return hostPortParts[0], "80"
		}
	}
	return "", ""
}

func getURI(url string) string {
	parts := strings.Split(url, ":")
	if len(parts) > 2 {
		uriParts := strings.Split(parts[2], "/")
		if len(uriParts) > 1 {
			return "/" + strings.Join(uriParts[1:], "/")
		}
	}
	return "/"
}

func Fetch(req HttpRequest, connection net.Conn) (HttpResponse, []Student, HttpRequest) {
	//This program handles the request-making to the server
	var res HttpResponse
	var students []Student

	requestBytes := RequestEncoder(req)

	connection.Write(requestBytes)

	responseBytes := make([]byte, BUFFER_SIZE)

	connection.Read(responseBytes)

	res = ResponseDecoder(responseBytes)
	if res.ContentType == "application/xml" {
		var xmlStudents Students
		err := xml.Unmarshal([]byte(res.Data), &xmlStudents)
		if err != nil {
			log.Fatalln(err)
		}
		trimmedXML := strings.TrimPrefix(res.Data, "<Students>")
		trimmedXML = strings.ReplaceAll(trimmedXML, "</Students>", "")
		res.Data = trimmedXML
		students = xmlStudents.Student
	} else if res.ContentType == "application/json" {
		filteredData := strings.ReplaceAll(res.Data, "\x00", "")
		err := json.Unmarshal([]byte(filteredData), &students)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return res, students, req

}

func ResponseDecoder(bytestream []byte) HttpResponse {
	var res HttpResponse

	responseStr := string(bytestream)

	headerAndData := strings.Split(responseStr, "\r\n\r\n")
	if len(headerAndData) > 1 {
		header := headerAndData[0]
		data := headerAndData[1]

		headerLines := strings.Split(header, "\r\n")

		responseLineParts := strings.Fields(headerLines[0])
		if len(responseLineParts) >= 2 {
			res.Version = responseLineParts[0]
			res.StatusCode = responseLineParts[1]
		}

		for _, line := range headerLines[1:] {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				switch parts[0] {
				case "Content-Type":
					res.ContentType = parts[1]
				case "Content-Language":
					res.ContentLanguage = parts[1]
				}
			}
		}
		res.Data = data
	}

	return res

}

func RequestEncoder(req HttpRequest) []byte {
	var result string
	requestLine := fmt.Sprintf("%s %s %s", req.Method, req.Uri, req.Version)

	headers := fmt.Sprintf("Host: %s\r\nAccept: %s\r\nAccept-Language: %s\r\n", req.Host, req.Accept, req.AcceptLanguage)

	result = fmt.Sprintf("%s\r\n%s\r\n", requestLine, headers)

	return []byte(result)

}
