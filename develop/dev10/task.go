package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Telnet utility ===

Implement a primitive telnet client:
Examples of calls:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

The program must connect to the specified host (ip or domain name) and port using the TCP protocol.
After connecting, the STDIN of the program must be written to the socket, and the data received from the socket must be output to STDOUT
Optionally, you can pass a timeout for connecting to the server to the program (via the --timeout argument, default 10s).

When you press Ctrl+D, the program should close the socket and exit. If the socket is closed on the server side, the program must also exit.
When connecting to a non-existent server, the program must terminate after a timeout.
*/

func main() {
	// Define command line arguments
	timeout := flag.Duration("t", 10*time.Second, "connection timeout")
	flag.Parse()

	// Validate arguments
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Args()[0]
	port := flag.Args()[1]
	address := net.JoinHostPort(host, port)

	// Connect to the specified address with a timeout
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	// Create channels for communication
	done := make(chan struct{})

	// Read from the server and write to STDOUT
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			fmt.Println("Error reading from connection:", err)
		}
		done <- struct{}{}
	}()

	// Read from STDIN and write to the server
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Println("Error writing to connection:", err)
		}
		done <- struct{}{}
	}()

	// Wait for either read or write to complete
	<-done
	fmt.Println("\nConnection closed")
}
