package main

import (
	"fmt"
)

func init() {
	fmt.Println("init running") //*先执行init函数
}

func main() {
	fmt.Println("main running")

	s := []int{1, 2, 3}
	for i := range s {
		if i == 0 {
			s[i] = 22
		}
	}
	fmt.Println(s)
}
