package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 15; i++ {
		ch <- i + 1 //*发送数据到channel
		time.Sleep(time.Second)
	}
	//!channel关闭，所有等待goroutine接收数据操作都将返回
	//!consume读取为空时会panic:deadlock
	close(ch)
}

func consume(ch <-chan int) {
	//*从channel接收数据，若没有数据则挂起直到channel关闭，否则panic:deadlock
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	ch := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		produce(ch)
	}()

	go func() {
		defer wg.Done()
		consume(ch)
	}()

	wg.Wait()
}
