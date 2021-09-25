package hanoi

import "fmt"

func Hanoi(n int, a, b, c string) {
	if n == 1 {
		move(a, c)
	} else {
		Hanoi(n-1, a, c, b) //~将n-1个盘子由A经过C移动到B
		move(a, c)          //&执行最大盘子移动
		Hanoi(n-1, b, a, c) //?剩下盘子由B经过A移动到C
	}
}

func move(a, b string) { //*执行最大盘子n从A到C
	fmt.Printf("move %s-->%s\n", a, b)
}
