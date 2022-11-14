package main

import (
	"log"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			active <- struct{}{}
			log.Printf("handle job:%d\n", j)
			time.Sleep(time.Second)
			<-active
		}(j)
	}

	wg.Wait()
}
