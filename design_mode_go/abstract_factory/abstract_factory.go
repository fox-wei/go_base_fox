package abstract_factory

import (
	"fmt"
)

//*抽象产品
type Lunch interface {
	Cook()
}

type Rice struct{}

func (r *Rice) Cook() {
	fmt.Println("I'm cooking the rice")
}

type Vegetable struct{}

func (v *Vegetable) Cook() {
	fmt.Println("Cooking the vegetable, the taste is good")
}

//*抽象工厂
type lunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

//?具体工厂
type SimpleLunchFactory struct{}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rice{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Vegetable{}
}

func NewlunchFactory() lunchFactory {
	return &SimpleLunchFactory{}
}
