package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var res = 0

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			res += 1
		}()
	}
	fmt.Println(res)
	wg.Wait()
}
