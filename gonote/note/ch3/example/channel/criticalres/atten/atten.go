package main

import (
	"fmt"
	"sync"
	"time"
)

var critical int = 0

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go func(mut *sync.Mutex) { //*采用指针，同一对象
		defer wg.Done()
		defer mut.Unlock()
		mut.Lock()
		critical = 23
		time.Sleep(time.Second * 10)
		fmt.Printf("g1:%d\n", critical)
	}(&mutex)

	time.Sleep(time.Second)
	mutex.Lock()
	critical = 24
	fmt.Printf("g2:%d\n", critical)
	mutex.Unlock()

	wg.Wait()
}
