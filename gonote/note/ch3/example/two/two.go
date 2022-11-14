package main

import (
	"fmt"
)

type Func func(string)

func Hello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

//*装饰器
func MyDecorator(fn Func) Func {
	return func(s string) {
		fmt.Println("好久不见了，我想你!")
		fn(s)
	}
}

func main() {
	s := MyDecorator(Hello)
	s("ying")
}
