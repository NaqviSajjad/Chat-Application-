package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {

	fmt.Println("Hello this line")
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		result := strconv.Itoa(rand.Intn(100)) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}

func main() {

	arguments := os.Args
	var PORT string

	if len(arguments) != 1 {
		for i, a := range os.Args[1:] {
			fmt.Printf("Argument %d is %s\n", i+1, a)
		}
		// fmt.Println("Please provide a port number!")
		// return
	}

	// fmt.Scan(&PORT)

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {

		c, err := l.Accept()
		if err != nil {

			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
