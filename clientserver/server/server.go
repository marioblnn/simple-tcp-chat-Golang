package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	port := ":8080"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Server error: ", err)
		os.Exit(1)
	}

	defer ln.Close()
	fmt.Println("Server listening on port: ", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Could not connect: ", err)
			continue
		}
		
		go handleConnection(conn)
		}
}

	

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s \n", clientAddr)

	conn.Write([]byte("Welcome\n"))
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected", clientAddr)
			break
		}

		message = strings.TrimSpace(message)
		fmt.Printf("%s %s", clientAddr, message + "\n")

		conn.Write([]byte(message + "\n"))

		if strings.ToLower(message) == "exit" {
			fmt.Printf("Client %s disconnected", clientAddr)
			break
		}
	}

}
