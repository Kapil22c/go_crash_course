package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sumOf(10, 8))
	total, err := sumOf(10, 5)
	if err != nil {
		fmt.Println("There was an error")
	} else {
		fmt.Println(total)
	}
}

func sumOf(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("start is great than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
