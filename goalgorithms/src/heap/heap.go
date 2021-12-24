package heap

//*自底向上实现大顶堆
func HeapBottomUp(arr []int) []int {
	for i := (len(arr) - 1) / 2; i >= 1; i-- {
		k := i
		v := arr[k]
		heap := false
		for !heap && 2*k < len(arr) {
			j := 2 * k
			if j < len(arr) {
				if arr[j] < arr[j+1] {
					j++
				}
			}
			if v >= arr[j] {
				heap = true
			} else {
				arr[k] = arr[j]
				k = j
			}
		}
		arr[k] = v
	}

	return arr
}
