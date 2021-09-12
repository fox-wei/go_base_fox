//请在此处输入您的练习代码
package main

import "fmt"

func showRectangle(width, weigth int) {
	flag1 := true
	flag2 := true
	g1 := "*"
	g2 := "*"
	for i := 1; i <= weigth; i++ {
		if i == 1 || i == weigth {
			if flag1 {
				s := "*"
				for j := 1; j <= width; j++ {
					g1 += s
				}
				flag1 = false
			}
			fmt.Println(g1)
		} else {
			if flag2 {
				n1 := " "
				for j := 1; j <= width; j++ {
					if j == width {
						g2 += "*"
					}
					g2 += n1
				}
				flag2 = false
			}
			fmt.Println(g2)
		}
	}
}

func main() {
	showRectangle(20, 10)
}
