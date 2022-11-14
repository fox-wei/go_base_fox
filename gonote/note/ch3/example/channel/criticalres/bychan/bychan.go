package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c chan int
	i int //?计数器
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}

	//*创建一个goroutine进行自增
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()

	return cter
}

func (cter *counter) Increase() int {
	return <-cter.c //*获取i
}

func main() {
	cter := NewCounter()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v := cter.Increase()
			fmt.Printf("Goroutine~%d:current value:%d\n", i, v)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println(cter.i)
}
