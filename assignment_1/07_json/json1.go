package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("JSON tutorial")

	author := Author{Name: "Kapil bavisiya", Age: 21, Developer: true}

	book := Book{Title: "My Biography", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
