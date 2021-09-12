package binarysearch

import "fmt"

//*输入：有序的集合
//*每次查找排除一半

func BinarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0 //^记录查找次数

	for {
		step = step + 1
		if low <= high { //!容易出错
			mid := (low + high) / 2 //*从中间开始
			guess := list[mid]

			//*找到目标值
			if guess == target {
				fmt.Printf("共查找了%d次\n", step)
				return mid
			}

			//*目标值小于中间值
			if guess > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}
