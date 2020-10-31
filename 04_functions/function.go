package main

import "fmt"

func main() {
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("can't divide by zero")
	}
	return a / b, nil

	// anonymous funciton
	// func() {
	// 	fmt.Println("Hii Kapil")
	// }() ... that parameter invokes this function

	// for i := 0; i < 5; i++ {
	// 	func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	f := func() {
		fmt.Println("Hii Kapil")
	}
	f()
}
