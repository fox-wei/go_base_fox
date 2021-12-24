package dynamic

import "fmt"

//*一般动态规划解决普通0-1背包问题
//^返回重量
func Knapsack1(weight []int, n, w int) int {
	states := make([][]bool, n)
	s1 := make([]bool, w+1)
	for k := range states {
		states[k] = s1
	}
	//?第一个物品初始化
	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ {
		//*前i-1个背包的信息;i-1行信息更新到i行
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}

		//*更新第i个物品信息；更新i行信息
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][weight[i]+j] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}

//?动态规划优化
func Knapsack2(items []int, n, w int) int {
	states := make([]bool, w+1)
	states[0] = true
	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := w - items[i]; j >= 0; j-- {
			if states[j] {
				states[items[i]+j] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

//!背包问题-计算最大价值
func Knapsack3(weight []int, values []int, n, w int) int {
	states := make([][]int, n) //*保存价值
	s1 := make([]int, w+1)
	for k := range states {
		states[k] = append(states[k], s1...)
	}

	// for j := range states {
	// 	for k := range states {
	// 		states[j][k] = -1
	// 	}
	// }

	//?初始化
	// states[0][0] = 0
	// fmt.Println(states[0])
	if weight[0] <= w {
		states[0][weight[0]] = values[0]
	}
	fmt.Println(states[0])
	//?动态规划实现
	for i := 1; i < n; i++ {
		//*前一个物品初始化
		for j := 0; j <= w; j++ {
			if states[i-1][j] > 0 {
				states[i][j] = states[i-1][j]
			}
		}

		//?放入第i个物品
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] > 0 {
				v := states[i-1][j] + values[i]
				if v > states[i][j+weight[i]] {
					states[i][j+weight[i]] = v
				}
			}
		}

		fmt.Println(states[i])
	}

	//*找最大值
	max := -1
	for i := 0; i <= w; i++ {
		if states[n-1][i] > max {
			max = states[n-1][i]
		}
	}
	return max
}
