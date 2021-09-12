package main

import "fmt"

type Slice struct {
	slice []int
	count int
}

//!数据添加到开头
func (s *Slice) Add(val int) {
	if s.count < cap(s.slice) {
		s.slice[s.count-1] = s.slice[0]
	} else {
		s.slice = append(s.slice, s.slice[0])
	}
	s.slice[0] = val
	s.count++
}

//!删除末尾
func (s *Slice) Delete() {
	s.count--
}

func main() {
	slice := &Slice{[]int{1}, 1}
	slice.Add(22)
	slice.Add(24)
	slice.Add(26)
	fmt.Println(slice.slice)
}
