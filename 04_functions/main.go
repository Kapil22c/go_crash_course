package main

import "fmt"

func greeting(name string) string {
	return "Hellooo " + name
}

func getSum(num1 int, num2 int) int {
	return num1+ num2
}
func main() {
	fmt.Println(greeting("Kapil"))
	fmt.Println(getSum(5,9))
}
