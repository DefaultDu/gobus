package main

import (
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte,1024)

    for {
        n, err := conn.Read(buffer)
        if err!= nil {
            fmt.Println("Error reading:", err)
            break
        }

        fmt.Println("Received:", string(buffer))
    }
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
    if err != nil {
        panic(err)
    }
	defer listener.Close()

	fmt.Println("Listening on port 8888")

    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }

        go handleConnection(conn)
    }
}