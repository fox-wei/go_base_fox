package main

import "fmt"

//*利用分治策略实现数组求和

func main() {
	total := sum([]int{1, 3, 5, 7, 9})
	fmt.Println(total)
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + sum(arr[1:])
}
