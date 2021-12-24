package heap

import (
	"fmt"
	"testing"
)

func TestHeapBottomUp(t *testing.T) {
	arr := []int{0, 1, 8, 6, 5, 3, 7, 4}
	fmt.Println(HeapBottomUp(arr))
}
