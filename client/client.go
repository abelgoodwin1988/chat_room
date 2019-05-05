package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

// Handle closing channel as a result of message from server
var closeChan = make(chan bool, 1)

func main() {
	conn, err := net.Dial("tcp", ":23")
	if err != nil {
		log.Fatalf("Failed to dial to port 23.\nError: %v\n", err)
	}
	// cleanup and close
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	go func(conn net.Conn) {
		for range interruptChan {
			log.Printf("Closing Client")
			conn.Write([]byte("\\close"))
			conn.Close()
			os.Exit(0)
		}
	}(conn)
	defer conn.Close()

	log.Printf("Connected to server at: %v", conn.RemoteAddr())
	go send(conn)
	go receive(conn)
	// Blocker/close channel
	for range closeChan {
		log.Printf("Closing client...\n")
		conn.Write([]byte("\\close"))
		conn.Close()
		os.Exit(0)
	}
}

func send(conn net.Conn) {
	for {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(text))
	}
}
func receive(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// closeChan trigger if message from server indicates
		if message == "\\close" {
			log.Printf("Received close signal from server...\n")
			closeChan <- true
		}
		fmt.Printf("%v", message)
	}
}
