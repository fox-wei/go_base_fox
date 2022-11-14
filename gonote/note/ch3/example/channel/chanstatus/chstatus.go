package main

import (
	"fmt"
	"sync"
	"time"
)

//*对channel判空和读取同时完成
//*使用select

func trySend(ch chan<- int, i int) bool {
	select {
	case ch <- i:
		return true
	default:
		return false
	}
}

func tryRecv(ch <-chan int) (int, bool) {
	select {
	case i := <-ch:
		return i, true
	default:
		return 0, false
	}
}

//^生产者
func producer(ch chan<- int) {
	var i int = 0
	for {
		if i > 10 {
			return
		}

		time.Sleep(time.Second * 2)

		ok := trySend(ch, i)
		if ok {
			fmt.Printf("[producer] send [%d] to channel\n", i)
			i++
			continue
		}

		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

//^消费者
func consumer(ch <-chan int) {
	for {
		i, ok := tryRecv(ch)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("[consumer]: recv [%d] from channel\n", i)
		if i > 6 {
			fmt.Println("[consumer]: exit......")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(ch)
	}()

	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()

}
