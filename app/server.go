package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

    var conn net.Conn
	conn, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
    defer conn.Close()

    reader := bufio.NewReader(conn)
    buffer := make([]byte, 1024)

    var n int
    n, err = reader.Read(buffer)
    if err != nil {
		fmt.Println(err)
        os.Exit(1)
    }

    http_request := strings.Split(string(buffer[:n]), "\r\n")
    start_line := http_request[0]
    fields := strings.Fields(start_line)
    if len(fields) < 3 {
        fmt.Println("invalid start line")
    }

    var target string
    _, target, _ = fields[0], fields[1], fields[2]
    respond_with_content := len(strings.Split(target, "/echo/")) > 1

    var status_line string
    if target == "/" || respond_with_content {
        status_line = "HTTP/1.1 200 OK\r\n"
    } else {
        status_line = "HTTP/1.1 404 Not Found\r\n"
    }

    var response_body string
    if respond_with_content {
        response_body = strings.Split(target, "/")[2]
    } else {
        response_body = ""
    }

    content_length := len(response_body)
    headers := []string{
        "Content-Type: text/plain",
        fmt.Sprintf("Content-Length: %d", content_length),
    }

    header_string := strings.Join(headers, "\r\n")
    response := fmt.Sprintf("%s\r\n%s\r\n\r\n%s", status_line, header_string, response_body)

    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Error writing response:", err)
    }

    conn.Close()
}

