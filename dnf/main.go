package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	flagPort = flag.Int("port", 53, "Port for the DNS forwarder server")
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("error running: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	flag.Parse()

	// Start a UDP server
	addr := fmt.Sprintf("127.0.0.1:%d", *flagPort)
	srvAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("resolving UDP address %s: %w", addr, err)
	}

	// Create a UDP listener.
	conn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		return fmt.Errorf("listening: %w", err)
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on", srvAddr)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("failed to read from UDP:", err)
			continue
		}

		// Display received data
		fmt.Printf("received data from %v: %s\n", clientAddr, string(buffer[:n]))

		// Send a response back to the client
		response := []byte("Hello from UDP server!")
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
