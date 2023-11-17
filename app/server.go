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

    _, target, _ := fields[0], fields[1], fields[2]

    var response string
    if target == "/" {
        response = "HTTP/1.1 200 OK\r\n\r\n"
    } else {
        response= "HTTP/1.1 404 Not Found\r\n\r\n"
    }

    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Error writing response:", err)
    }

    conn.Close()
}

