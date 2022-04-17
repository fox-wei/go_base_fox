package main

import (
	"fmt"
	"reflect"
)

func dumpMethodSet(i interface{}) {
	dyType := reflect.TypeOf(i)

	if dyType == nil {
		fmt.Printf("There is no dynamic type\n")
		return
	}

	n := dyType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method is empty!\n", dyType)
		return
	}

	fmt.Printf("%s's  method set is:\n", dyType)
	for j := 0; j < n; j++ {
		fmt.Println("-", dyType.Method(j).Name)
	}
}

type T struct{}

func (T) M1() {}
func (T) M2() {}

type S T

func main() {
	var s S
	dumpMethodSet(&s)
}
