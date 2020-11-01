package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	// for j := 0; j < 5; j++ {
	wg.Add(2)

	// recieving data only channel
	go func(ch <-chan int) {
		// i := <-ch
		// fmt.Println(i)
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		// i = <-ch
		// fmt.Println(i)
		// ch <- 27 .. can't do it as it is receive only channel
		wg.Done()
	}(ch)

	// send data only channel
	go func(ch chan<- int) {
		// i := 42
		// ch <- i
		// i = 27
		ch <- 42
		// close(ch) // this will lead to panic on channel
		ch <- 40
		close(ch) // for avoid deadlock in channel

		// fmt.Println(<-ch)
		wg.Done()
	}(ch)

	// }
	wg.Wait()
}
