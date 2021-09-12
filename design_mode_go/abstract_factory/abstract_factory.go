package abstract_factory

import (
	"fmt"
)

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

type lunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

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
