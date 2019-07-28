package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Print("Error in connection")
		return
	} else {
		fmt.Print("Connection succesfull \n")
	}
	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received from client:", string(message))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		newmessage, _ := reader.ReadString('\n')
		fmt.Scanln(newmessage)

		// sample process for string received
		//newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
