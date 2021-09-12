package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{22, 23, 24, 26, 1301, 1602, 1901, 1902} //?有序
	fmt.Println("寻找值22的位置:")
	res := BinarySearch(s, 22)
	fmt.Printf("所在位置为%d\n", res)
}
