package selectsort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	s := []int{1301, 24, 22, 26, 23, 1601, 1902, 1901}
	fmt.Printf("未排序的：%v\n", s)

	s = SelectionSort(s)
	fmt.Printf("升序排序...\n%v\n", s)
}
