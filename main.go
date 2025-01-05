package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":5555"
	listner, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("ERROR: Could not initizalize server! - ", err)
		os.Exit(1)
	}
	defer listner.Close()

	fmt.Println("Server waiting for connection at port " + port)

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("ERROR: Connection not accepted! - ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected: " + clientAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(clientAddr + ": " + msg)

		_, err := conn.Write([]byte("Message Recieved!\n"))
		if err != nil {
			fmt.Println("ERROR: Message not sent! - ", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR: Connection with client refused! - ", err)
	}

	fmt.Println("Connection with " + clientAddr + " closed")
}
