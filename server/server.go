// package main contains client code for a multi-client tcp chat server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

// Client contains a clients name and net.Conn information
type Client struct {
	name string
	conn net.Conn
}

var clients []Client                    // maintains all current clients
var messages = make(chan string)        // communicates client messages
var closeChan = make(chan os.Signal, 1) // communicates closing of server

// main creates a listener and creates a loop for accepting
//	client connections, and spinning of handlers for that
//	client connection.
func main() {
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		log.Fatalf("Failed to open chat server at: %v\n", listener.Addr())
		listener.Close()
	}
	log.Printf("Opened chat server at: %v\n", listener.Addr())

	// cleanup and signal close
	signal.Notify(closeChan, os.Interrupt)
	go func(listener net.Listener) {
		for range closeChan {
			close(listener)
		}
	}(listener)
	defer listener.Close()

	// Create infinite loop for accepting all new client connections
	//	TODO: make name-setting non-blocking
	for {
		conn, err := listener.Accept()
		conn.Write([]byte("Enter a username:\n"))
		name, _ := bufio.NewReader(conn).ReadString('\n')
		newClient := Client{name: name[:len(name)-1], conn: conn}
		if err != nil {
			log.Printf("Failed to open connection:\nError: %v\n", err)
		}
		clients = append(clients, newClient)
		log.Printf("Added connection to pool at: %v\n", conn.RemoteAddr())
		go clientHandler(newClient)
	}
}

// clientHandler handles spinning off the concurrent reading of a newly
//	created client, and the writing to clients of any client messages
func clientHandler(client Client) {
	go readClientMessages(client)
	go writeClientMessages()
}

// readClientMessages creates a go routine per created client
//	which will read any new messages into the messages channel.
// 	Watched for a close command which closes and removes the
//	client from the clients slice and exits the go routine
func readClientMessages(client Client) {
	go func() {
		for {
			input, _ := bufio.NewReader(client.conn).ReadString('\n')
			log.Printf("Received Message from -:%v:- * %v", client.conn.RemoteAddr(), input)
			if input == "\\close" {
				log.Printf("Closing connection with client %v at %v", client.name, client.conn.RemoteAddr())
				closeClientConnection(client)
				return
			}
			t := time.Now()
			messages <- fmt.Sprintf("%q %q: %q", t.Format(time.RFC822), client.name, input)
		}
	}()
}

// writeClientMessages ranges over all new messages passed to the
//	messages channel. It iterates over our slice of clients
//	and writes the message to all clients
func writeClientMessages() {
	for message := range messages {
		log.Printf("Writing Message: %v", message)
		for _, client := range clients {
			client.conn.Write([]byte(message))
		}
	}
}

// closeCientConnection takes in a client, fines that client
//	in the global clients slice, and removes it.
func closeClientConnection(client Client) {
	for i, c := range clients {
		if c.conn.RemoteAddr() == client.conn.RemoteAddr() {
			clients = append(clients[:i], clients[i+1:]...)
		}
	}
}

func close(listener net.Listener) {
	// Close all current connections
	for _, client := range clients {
		log.Printf("Closing Client %v", client.name)
		closeClientConnection(client)
		client.conn.Write([]byte("\\close"))
		client.conn.Close()
	}
	// Close listner, so as to accept no new connections
	listener.Close()
	os.Exit(0)
}
