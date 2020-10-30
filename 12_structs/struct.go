package main

import "fmt"

type Doctor struct {
	Number     int
	ActorName  string
	Episodes   []string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Dr. Sanket Bhosle",
		Companions: []string{
			"Liz shaw",
			"Jo Grant",
			"Nikhil Mehta",
		},
	}
	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.Companions)

	// anonymous struct
	// aDoctor := struct{name string} {name: "Kapil Patel"}
	// anotherDoctor := aDoctor // if you put &before aDoctor then it will be pointer
	// so it will not print both name.. it will print only one name
	// anotherDoctor.name = "Tom Steyn"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)
	// it will print both the outputs

}
