package selectsort

//*选择排序
//*每次选出最大或最小的数

//^寻找最小的数
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_indext := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_indext = i
		}
	}
	return smallest_indext
}

//^排序函数
func SelectionSort(arr []int) []int {
	result := []int{}

	count := len(arr)
	for i := 0; i < count; i++ {
		smallest_index := findSmallest(arr)
		result = append(result, arr[smallest_index])
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...) //!易出错，+1
	}

	return result
}

//^Bubble sort，升序
func BubbleSort(arr []int) []int {
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i] //?两两交换
				flag = true                         //*如果不交换则排序完成
			}
		}
	}
	return arr
}
