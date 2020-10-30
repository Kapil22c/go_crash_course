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

	// another example
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Surat":     4523020,
		"Ahmedabad": 7520125,
		"Rajkot":    2024551,
		"Mumbai":    125800212,
		"Chennai":   2510212,
	}
	// array is valid key type but slice is not valid for maps
	// m := map[[0]int]string{}
	// fmt.Println(statePopulations, m)

	// built in make function
	// fmt.Println(statePopulations)

	pop, ok := statePopulations["Surat"]
	fmt.Println(pop, ok)
	// ok will return true or false as per it finds key or not

	fmt.Println(len(statePopulations))
}
