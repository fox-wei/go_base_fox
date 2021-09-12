package facade

import (
	"fmt"
	"testing"
)

func TestNewCarFacade(t *testing.T) {
	f := NewCarFacade()
	f.CreateCompleteCar()

	fmt.Println(DecBase(023))
	fmt.Println(DecBase(0x12))
	fmt.Println(DecBase(1010))
}
