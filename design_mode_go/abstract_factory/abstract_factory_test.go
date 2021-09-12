package abstract_factory

import (
	"testing"
)

func TestNewlunchFactory(t *testing.T) {
	factory := NewlunchFactory()

	rice := factory.CreateFood()
	rice.Cook()

	vegetalbe := factory.CreateVegetable()
	vegetalbe.Cook()
}
