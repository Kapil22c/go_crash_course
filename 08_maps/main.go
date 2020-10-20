package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	//Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Kapil"] = "kapil@gmail.com"
	// emails["KK"] = "KK@gmail.com"

	// DEclare map and add kv
	emails := map[string]string{"Bob": "Bob@gmail.com",
		"Kapil": "kapil@gmail.com",
		"KK":    "Kk@gmail.com"}

	emails["nehal"] = "nehal@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Kapil"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
