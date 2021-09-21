package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := []int{22, 23, 12, 34, 56, 12, 56, 10, 2, 45, 23, 26}
	fmt.Println(QuickSort(s))
}
