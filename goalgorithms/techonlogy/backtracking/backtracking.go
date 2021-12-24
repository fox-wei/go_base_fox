package backtracking

var maxW int //&最大重量
var weight = []int{2, 2, 4, 6, 3}
var n = 5  //*物品数量
var w = 11 //^最大承受重量

//*回溯法解决0-1背包问题-一般回溯法
func F1(i, cw int) {
	if cw == n || n == i {
		if cw > maxW {
			maxW = cw
		}
		return
	}

	F1(i+1, cw)
	if cw+weight[i] <= w {
		F1(i+1, cw+weight[i])
	}
}

//&使用备忘录进行优化
var mem [5][10]bool //?备忘录
func F2(i, cw int) {
	if cw == n || n == i {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	//*重复状态
	if mem[i][cw] {
		return
	}
	mem[i][cw] = true
	F1(i+1, cw)
	if cw+weight[i] <= w {
		F1(i+1, cw+weight[i])
	}
}
