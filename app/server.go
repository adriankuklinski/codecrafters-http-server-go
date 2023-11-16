package main

import (
    "bufio"
	"fmt"
	"net"
	"os"
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
    _, err = reader.Read(buffer)
    if err != nil {
		fmt.Println(err)
        os.Exit(1)
    }

    response := "HTTP/1.1 200 OK\r\n\r\n"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Error writing response:", err)
    }

    conn.Close()
}

