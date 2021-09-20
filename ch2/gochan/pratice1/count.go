package main

import (
	"fmt"
	"sync"
)

func goroutine(name string, share chan int) {
	for {
		value, ok := <-share
		if !ok {
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		share <- (value + 1)
	}
}

func main() {
	share := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		goroutine("ying", share)
	}()

	go func() {
		defer wg.Done()
		goroutine("yi", share)
	}()

	share <- 1

	wg.Wait()
}
