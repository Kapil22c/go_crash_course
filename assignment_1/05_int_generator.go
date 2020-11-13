package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var total = 10000
	ch := make(chan int, total)
	single := make([]int, total, total)
	fmt.Printf("Initiating random number generation.\n")
	start := time.Now()
	go makeRandomNumbers(total, ch)
	for i := 0; i < total; i++ {
		single = append(single, (<-ch))
	}
	elapsedSingleRun := time.Since(start)
	fmt.Printf("Run took %s\n", elapsedSingleRun)
}

func makeRandomNumbers(numInts int, ch chan int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < numInts; i++ {
		ch <- generator.Intn(numInts)
	}
	close(ch)

	for item := range ch {
		fmt.Println(item)
	}
}
