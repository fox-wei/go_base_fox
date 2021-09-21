package quicksort

/*
*1. find the pivot
*2. divide the slice by pivot
*3. recurse the quick sort
^The best case:nlogn and the worse case:n^2
*/

func QuickSort(arr []int) []int {
	if len(arr) < 2 { //?Base Case
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}

	return append(QuickSort(left),
		append([]int{pivot}, QuickSort(right)...)...)
}
