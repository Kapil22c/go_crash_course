package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide host:Port")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send Message -> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		//message, _ := bufio.NewReader(c).ReadString('\n')
		//fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "stop" {
			fmt.Println("Stopping TCP server")
			fmt.Println("-------------------------")
			return
		}
		if strings.TrimSpace(string(text)) == "printAll" {
			fmt.Println("All data")
			fmt.Println("-------------------------")

		}

		if strings.TrimSpace(string(text)) == "new" {
			fmt.Print("Username: ")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "break" {
			fmt.Print("You want to break communication? : ")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "first" {
			fmt.Println("First message of username: ")
			fmt.Println("-------------------------")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "print" {
			fmt.Println("Printing data")
			fmt.Println("-------------------------")

		}

	}
}
