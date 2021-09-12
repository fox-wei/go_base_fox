package main

import "fmt"

func FiboaccaiByswap(start, end int) []int {
	a := 0
	b := 1
	FiboaccaiSlice := []int{}
	for a < end {
		if a >= start {
			FiboaccaiSlice = append(FiboaccaiSlice, a)
		}
		a, b = b, a+b
	}
	return FiboaccaiSlice
}

func main() {
	s := FiboaccaiByswap(0, 100)
	fmt.Println(s)
}
