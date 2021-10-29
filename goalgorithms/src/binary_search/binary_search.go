package binarysearch

import "fmt"

//*输入：有序的集合
//*每次查找排除一半

func BinarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0 //^记录查找次数

	for low <= high {
		step = step + 1

		mid := (low + high) / 2 //*从中间开始
		guess := list[mid]

		//*找到目标值
		if guess == target {
			fmt.Printf("在%d找到,共查找了%d次\n", mid, step)
			return mid
		}

		//*目标值小于中间值
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Printf("%d找不到，共查找了%d次\n", target, step)
	return -1
}

//^三分查找
func TrisectionSearch(arr []int, val int) (int, int) {
	low := 0
	high := len(arr) - 1

	step := 0 //^记录查找次数

	for low <= high {
		step += 1
		mid1 := low + int((high-low)/3)
		mid2 := high - int((high-low)/3)

		// 枢轴
		midData1 := arr[mid1]
		midData2 := arr[mid2]

		if midData1 == val {
			return mid1, step
		} else if midData2 == val {
			return mid2, step
		}

		if midData1 < val {
			low = mid1 + 1
		} else if midData2 > val {
			high = mid2 - 1
		} else {
			low = low + 1
			high = high - 1
		}
	}
	return -1, step
}

//^插值插入
func MidSearch(arr []int, val int) (int, int) {

	low := 0
	high := len(arr) - 1

	step := 0
	//循环的终止条件
	for low <= high {
		step += 1

		// 大段
		leftv := float64(val - arr[low])

		// 整段
		allv := float64(arr[high] - arr[low])

		// 整段差
		diff := float64(high - low)

		// 计算中间值
		mid := int(float64(low) + diff*leftv/allv)

		if mid < 0 || mid >= len(arr) {
			return -1, step
		}

		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			return mid, step
		}
	}
	return -1, step

}
