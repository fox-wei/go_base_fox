package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	fmt.Println("Worker is working...")
	time.Sleep(time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		fmt.Println("Worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a worker...")
	c := spawn(worker)
	<-c //!等待goroutine退出
	fmt.Println("Worker work done!")
}
