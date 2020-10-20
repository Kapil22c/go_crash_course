package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign at same time
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])

	// Array slice 
	fruitSlice := []string{"Apple","Mango", "Grapes", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))	
	fmt.Println(fruitSlice[1:3])	

}
