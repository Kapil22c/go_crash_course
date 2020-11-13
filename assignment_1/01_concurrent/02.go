package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"paneer tikka", "butter chicken", "pasta", "cake"}
	results := make(chan bool)
	// results <- true
	// <- result
	for _, f := range foods {
		go func(f string) {
			cook(f)
			results <- true
		}(f)
	}
	for i := 0; i < len(foods); i++ {
		<-results
	}
	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
