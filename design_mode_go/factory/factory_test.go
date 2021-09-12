package factory

import (
	"testing"
)

func TestNewEatFool(t *testing.T) {
	NewEatFool("c").MyFavoriteFood()
	NewEatFool("a").MyFavoriteFood()
}
