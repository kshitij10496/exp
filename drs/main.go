package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("resolving: %v\n", err)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) != 2 {
		return errors.New("invalid syntax: needs exactly 2 arguments")
	}

	dnsServer, address := args[0], args[1]
	ipAddress, err := resolve(address, dnsServer)
	if err != nil {
		return fmt.Errorf("resolving address %s: %w", address, err)
	}
	fmt.Println(ipAddress)
	return nil
}
