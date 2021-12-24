package backtracking

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	F2(0, 0)
	fmt.Println(maxW)

	F1(0, 0)
	fmt.Println(maxW)
}
