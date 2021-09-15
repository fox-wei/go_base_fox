package main

import "fmt"

//*利用分治策略实现数组求和

func main() {
	total := sum([]int{1, 3, 5, 7, 9})
	fmt.Println(total)
	fmt.Println(goc(31415, 14142))
}

func sum(arr []int) int {
	if len(arr) == 0 {
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
