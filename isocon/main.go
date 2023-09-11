package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Welcome message
	conn.Write([]byte("Welcome to MyTelnetServer!\r\n"))

	// Read and write data between the client and the server
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
			return
		}

		// Handle EOF (Ctrl + D) signal from client.
		// Server should gracefully close the client connection.
		if n == 1 {
			if buffer[0] == 4 {
				fmt.Printf("closing client connection: %v\n", conn.RemoteAddr())
				break
			}
		}

		input := string(buffer[:n])
		fmt.Printf("Received from %s: %s", conn.RemoteAddr(), input)

		// Echo back to the client
		conn.Write([]byte(input))
	}
}

func main() {
	// Listen on port 23 (Telnet default port)
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Telnet server is listening on port 23...")

	// Handle incoming connections in a separate goroutine
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleClient(conn)
		}
	}()

	// Gracefully handle Ctrl+C to exit the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Server is shutting down...")
}
