package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//var temp string

func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func main() {
	q := make(map[string][]string, 10)

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("PORT number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	var queue []string
	var username string
	for {
		newdata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(newdata)) == "stop" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("Received message -> ", string(newdata))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))

		if strings.TrimSpace(string(newdata)) == "new" {
			fmt.Println("New user")

			//username := string(newdata)
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			username = string(newdata)
			q[username] = queue

		} else {
			queue = append(queue, string(newdata))
			//queue = enqueue(queue, string(newdata))
			q[username] = queue
		}

		if strings.TrimSpace(string(newdata)) == "break" {
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Command:", string(newdata))
			queue = nil
			//queue = enqueue(queue, string(newdata))
			//
		}

		if strings.TrimSpace(string(newdata)) == "print" {
			//fmt.Println("Queue:", q)
			for key, value := range q {
				value = value[:len(value)-1]
				fmt.Printf("%s value is %v\n", key, value)

			}
			//queue = enqueue(queue, string(newdata))
			//
		}

		if strings.TrimSpace(string(newdata)) == "first" {
			//fmt.Println("Queue:", q)
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			username = string(newdata)

			for key, value := range q {
				if key == username {
					fmt.Println("First message of ", key, " is ", value[:1])
				}

			}
			//queue = enqueue(queue, string(newdata))
			//
		}

		}
	}
}
