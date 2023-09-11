package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func (d *DBMS) handleClient(conn net.Conn) {
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

		cmd, err := parseCmd(input)
		if err != nil {
			fmt.Printf("parsing command from %s: %v", conn.RemoteAddr(), err)
			conn.Write([]byte(err.Error()))
			conn.Write([]byte("\r\n"))
			continue
		}

		switch cmd.Name {
		case CommandGet:
			val, err := d.db.Get(context.Background(), cmd.Args[0])
			if err != nil {
				if errors.Is(err, ErrNoKeyFound) {
					conn.Write([]byte(err.Error()))
					conn.Write([]byte("\r\n"))
					continue
				}
			}
			conn.Write([]byte(val))
			conn.Write([]byte("\r\n"))
		case CommandSet:
			if err := d.db.Set(context.TODO(), cmd.Args[0], cmd.Args[1]); err != nil {
				conn.Write([]byte(err.Error()))
				conn.Write([]byte("\r\n"))
				continue
			}
			conn.Write([]byte("saved"))
			conn.Write([]byte("\r\n"))
		}

		// // Echo back to the client
		// conn.Write([]byte(input))
	}
}

const (
	CommandGet = "get"
	CommandSet = "set"
)

var ErrNoKeyFound = errors.New("key not found")

type Cmd struct {
	Name string
	Args []string
}

func parseCmd(s string) (Cmd, error) {
	input := strings.Split(strings.ToLower(strings.TrimSpace(s)), " ")
	if len(input) == 0 {
		return Cmd{}, fmt.Errorf("missing command")
	}

	switch name := input[0]; name {
	case "get":
		if len(input) < 2 {
			return Cmd{}, fmt.Errorf("missing arguments from get command")
		}
		return Cmd{Name: CommandGet, Args: input[1:]}, nil
	case "set":
		if len(input) != 3 {
			return Cmd{}, fmt.Errorf("missing arguments for set command")
		}
		return Cmd{Name: CommandSet, Args: input[1:]}, nil
	default:
		return Cmd{}, fmt.Errorf("unsupported command: %s", name)
	}
}

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
}

type DBMS struct {
	listener net.Listener
	db       Store
}

func (d *DBMS) Serve() error {
	for {
		conn, err := d.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go d.handleClient(conn)
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

	db := DBMS{
		listener: listener,
		db: &ReadUncommittedDB{
			db: make(map[string]string),
		},
	}
	fmt.Println("DB server is listening on port 23...")

	// Handle incoming connections in a separate goroutine
	go db.Serve()

	// Gracefully handle Ctrl+C to exit the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Server is shutting down...")
}

type ReadUncommittedDB struct {
	db map[string]string
}

func (dbms *ReadUncommittedDB) Get(ctx context.Context, key string) (string, error) {
	if val, found := dbms.db[key]; found {
		return val, nil
	}
	return "", fmt.Errorf("%s: %w", key, ErrNoKeyFound)
}

func (dbms *ReadUncommittedDB) Set(ctx context.Context, key, value string) error {
	dbms.db[key] = value
	return nil
}
