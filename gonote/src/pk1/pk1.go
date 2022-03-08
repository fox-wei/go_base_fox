package main

import "fmt"

func init() {
	fmt.Println("init running") //*先执行init函数
}

func main() {
	fmt.Println("main running")
}
