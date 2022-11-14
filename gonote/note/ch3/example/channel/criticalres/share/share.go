package main

//!临界资源处理
//!传统的方式：共享内存+互斥锁

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	i int
}

var cter counter //~临界资源

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println(cter.i)
}
