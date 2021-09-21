package main

import "fmt"

//!Recusion:find the base case and recursion case
//*D&C: A strategy of recursion
/*
*step:
*1. find the base case
*2. divide the proble, recure the base case
 */

func main() {
	total := sum([]int{1, 3, 5, 7, 9})
	fmt.Println(total)
	fmt.Println(goc(31415, 14142))
}

//*利用分治策略实现数组求和
func sum(arr []int) int {
	if len(arr) == 0 { //&Base Case
		return 0
	}

	return arr[0] + sum(arr[1:])
}

//&Euclids's algorithm find the greatest common number
func goc(m, n int) int {
	for n != 0 {
		r := m % n
		m = n
		n = r
	}
	return m
}
