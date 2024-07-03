package main

import (
    "fmt"
    "net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	message := "Hello, server!"
    for {
        n, err := conn.Write([]byte(message))
        if err != nil {
            panic(err)
        }
        fmt.Printf("Message sent:", message)
    }
}