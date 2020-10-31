package main

import "fmt"

func main() {
	// sum("The sum is", 1, 5, 8, 6, 9)

	s := sum(1, 5, 8, 6, 4)
	fmt.Println("Answer is", s)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// func sum(values ...int) (result int) {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }


// in main fmt.Println("answer is", *s)

// func sum(values ...int) *int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result
// }
