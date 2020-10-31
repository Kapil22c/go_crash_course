package main

import "fmt"

func main() {
	// a := 42
	// b := &a
	// fmt.Println(a, *b)
	// a = 27
	// fmt.Println(a, *b)
	// *b = 14
	// fmt.Println(a, *b)
	// -------------------
	// a := [3]int{1, 4, 6}
	// b := &a[0]
	// c := &a[1]
	// fmt.Println("%v %p %p\n", a, b, c)
	// -------------------
	// var ms *myStruct
	// fmt.Println(ms)    // will return <nil>
	// ms = new(myStruct) // will give &{0}
	// fmt.Println(ms)
	// ms = &myStruct{foo: 43}
	// fmt.Println(ms)

	// ------------------
	greeting := "Hello"
	name := "Kapil"
	greet(&greeting, &name)
	fmt.Println(name)
}

func greet(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Kumar"
	fmt.Println(*name)
}

// type myStruct struct {
// 	foo int
// }
