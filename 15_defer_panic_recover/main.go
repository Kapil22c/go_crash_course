package main

import (
	"fmt"
	"log"
)

func main() {
	// defer..

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// a  := "start"
	// defer fmt.Println(a)
	// a = "end"
	//
	// this would print "start" as it defer takes
	//  args when it was assigned not after the fuction was called
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// Panic..
	// panic happens aftr defer statements are executed
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking")

}
