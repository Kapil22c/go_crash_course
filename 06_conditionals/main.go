package main

import "fmt"

func main() {
	x := 10
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is neither blue nor red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("nothing")
	}
}
