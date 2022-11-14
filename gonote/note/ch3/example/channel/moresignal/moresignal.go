package main

import (
	"fmt"
	"sync"
	"time"
)

//& 1:n信号通知机制
type signal struct{}

func worker(i int) {
	fmt.Printf("Worker %d: is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d: is done!\n", i)
}

func spawnGroup(f func(int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-groupSignal //~所有goroutine阻塞，直到channel关闭才能继续执行
			fmt.Printf("Worker %d:start to working...\n", i)
			f(i)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

func main() {
	fmt.Printf("Start a group of workers......\n\n")
	groupSignal := make(chan signal)

	c := spawnGroup(worker, 5, groupSignal) //^5个goroutine发生阻塞
	time.Sleep(time.Second * 5)

	fmt.Printf("The group of workers start to work...\n\n")
	close(groupSignal) //*发送信号，关闭channel让阻塞的goroutine开始工作，实现广播
	<-c                //?等待子goroutine退出
	fmt.Println()
	fmt.Printf("The group of workers work done!\n")
}
