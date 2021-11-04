package binarysearch

import (
	"fmt"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{22, 23, 24, 26, 1301, 1602, 1901, 1902} //?有序
	fmt.Printf("寻找值%d的位置:\n", 27)
	res := BinarySearch(s, 27)
	fmt.Printf("二分查找所在位置为%d\n", res)
	fmt.Printf("理论查找次数:%d\n", int(math.Log2(float64(len(s)))))
}
