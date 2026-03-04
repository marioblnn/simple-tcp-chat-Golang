package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) !=3 {
		fmt.Println("Use go run client.go <ip_server> <port>")
		os.Exit(1)
	}
	serverIp := os.Args[1]
	serverPort := os.Args[2]
	address := serverIp + ":" + serverPort

	conn, err := net.Dial("tcp", address)
	if err != nil{
		fmt.Println("Connection error: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to : ", address)
	
	
	serverReader := bufio.NewReader(conn)
	welcomeMsg, _ := serverReader.ReadString('\n')
	fmt.Println(welcomeMsg)

	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Send a message to the server: ")
		text, _ := input.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Fprintf(conn, "%s\n", text)

		if strings.ToLower(text) == "exit" {
			fmt.Println("Closing...")
			break
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection lost")
			break
		}
		fmt.Print("Sent to server: " + response)
	}
}