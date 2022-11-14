package main

import (
	"fmt"
	"time"

	"github.com/fox-wei/go_base_fox/gonote/note/ch3/example/workpool"
)

func main() {
	p := workpool.New(5)

	for i := 0; i < 5; i++ {
		err := p.Schedule(func() {
			time.Sleep(2 * time.Second)
		})

		if err != nil {
			fmt.Printf("task: %d, err: %v\n", i, err)
		}
	}

	time.Sleep(time.Second * 10)
	p.Free()
}
