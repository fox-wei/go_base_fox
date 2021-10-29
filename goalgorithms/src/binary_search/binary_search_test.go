package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{22, 23, 24, 26, 1301, 1602, 1901, 1902} //?有序
	fmt.Println("寻找值26的位置:")
	res := BinarySearch(s, 26)
	fmt.Printf("二分查找所在位置为%d\n", res)

	s1, s2 := TrisectionSearch(s, 26)
	fmt.Printf("三分查找，所在位置为%d，比较次数:%d\n", s1, s2)

	s3, s4 := MidSearch(s, 126)
	fmt.Printf("插值查找，所在位置为%d，比较次数:%d\n", s3, s4)
}
