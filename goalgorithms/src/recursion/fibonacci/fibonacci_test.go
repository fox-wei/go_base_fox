package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Fibonacci(i)
		fmt.Printf("%d\t", res)
	}

	fmt.Println()
}
