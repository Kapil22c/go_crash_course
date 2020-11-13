package main

import (
	"fmt"
)

type User struct {
	name    Name
	age     int
	city    string
	country string
}

type Name struct {
	firstName string
	lastName  string
}

func main() {
	user := User{
		Name{
			"Kapil",
			"Bavisiya",
		},
		21,
		"surat",
		"India",
	}
	fmt.Println("First Name:", user.name.firstName)
	fmt.Println("Last Name:", user.name.lastName)
	fmt.Println("Age:", user.age)
	fmt.Println("City:", user.city)
	fmt.Println("Country:", user.country)
}
