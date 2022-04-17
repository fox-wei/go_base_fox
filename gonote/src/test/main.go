package main

import "fmt"

func setup(name string) func() {
	fmt.Println("do some setup stuff for", name)
	return func() {
		//!闭包，使用局部变量name
		fmt.Println("do some teardown stuff for", name)
	}
}

func main() {
	t := setup("demo")
	defer t()
	fmt.Println("do some bussiness stuff")
}
