package main

import "fmt"

func main() {
	// var name = "kapil"
	// var age int64 = 21
	const isCool = true
	// isCool = false

	// Shorthand -- only can be use in function
	// name := "Kapil bavisiya"
	age := 21
	size := 1.3

	// email := "kapil@gmail.com"

	// we can do this as well
	name, email := "kapil", "kapil@gmail.com"
	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", isCool)
}
