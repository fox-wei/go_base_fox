package eratosthenes

import (
	"fmt"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	res := Eratosthenes(100)
	fmt.Println(res, len(res))
}
