package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("resolving: %v\n", err)
	}
}

func run() error {
	return fmt.Errorf("not implemented")
}
