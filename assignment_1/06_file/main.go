package main 

import ( 
	"fmt"
	"io/ioutil"
	"log"
	"os"
) 

func CreateFile() { 

	fmt.Printf("Writing to a file in Go lang\n") 

	file, err := os.Create("kloudOne.txt") 
	
	if err != nil { 
		log.Fatalf("failed creating file: %s", err) 
	} 

	defer file.Close() 
	
	len, err := file.WriteString("Welcome to KloudOne."+ 
			" This program demonstrates reading and writing"+ 
						" operations to a file in Go lang.") 

	if err != nil { 
		log.Fatalf("failed writing to file: %s", err) 
	} 
	fmt.Printf("\nFile Name: %s", file.Name()) 
	fmt.Printf("\nLength: %d bytes", len) 
} 

func ReadFile() { 

	fmt.Printf("\n\nReading a file in Go lang\n") 
	fileName := "test.txt"
	
	data, err := ioutil.ReadFile("kloudOne.txt") 
	if err != nil { 
		log.Panicf("failed reading data from file: %s", err) 
	} 
	fmt.Printf("\nFile Name: %s", fileName) 
	fmt.Printf("\nSize: %d bytes", len(data)) 
	fmt.Printf("\nData: %s", data) 

} 

func main() { 

	CreateFile() 
	ReadFile() 
} 
