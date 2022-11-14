package main

import "fmt"

func main() {
	// fmt.Println("Hello Go!")
	// s := "我的专业是软件工程"
	// fmt.Printf("len of the bytes:%d\n", len(s))
	// fmt.Printf("len of the char:%d\n", utf8.RuneCountInString(s))

	// r := &replaceHolder{"fox", 21}
	// fmt.Printf("v->%v\n", r)
	// fmt.Printf("+v->%+v\n", r)
	// fmt.Printf("#v->%#v\n", r)

	// m := []int{0, 1, 2, 3, 4, 5}
	// m1 := m[1:4]
	// fmt.Printf("m:%v\tm1:%v\n", m, m1)

	twoTimes := partitalTimes(2)
	fmt.Println(twoTimes(5))
	fmt.Println(twoTimes(6))
}

type replaceHolder struct {
	Name string
	Age  int
}

func times(x, y int) int {
	return x * y
}

func partitalTimes(x int) func(int) int {
	return func(i int) int {
		return times(x, i)
	}
}
