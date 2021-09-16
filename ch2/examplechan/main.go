package main

import (
	"fmt"
)

func unBuffter() {
	ch := make(chan int)

	var a string

	go func() {
		a = "hello world"
		<-ch
	}()

	ch <- 0

	fmt.Println(a)
}

func bufftered() {
	ch := make(chan int, 10)

	var a string

	go func() {
		a = "hello ying"
		ch <- 0
	}()

	<-ch

	fmt.Println(a)
}

func closed() {
	ch := make(chan int, 10)
	var a string

	go func() {
		a = "Hello qi"
		close(ch)
	}()

	<-ch

	fmt.Println(a)
}

func main() {
	unBuffter()
	bufftered()
	closed()
}
