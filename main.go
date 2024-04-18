// eth3r-network, 2023-2024
// based off https://dev.to/hgsgtk/how-go-handles-network-and-system-calls-when-tcp-server-1nbd

package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections.
	addr := "localhost:2142"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	host, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Handle connections in a new goroutine
		go func(conn net.Conn) {
			for {
				buf := make([]byte, 1024)
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("Error reading: %#v\n", err)
					return
				}
				fmt.Printf("Message received: %s\n", string(buf[:len]))
			}
		}(conn)
	}
}
