package factory

import "fmt"

//?定义接口，吃
type EatFool interface {
	MyFavoriteFood()
}

type Chinese struct{}

func (c Chinese) MyFavoriteFood() {
	fmt.Println("Chinese like eating the rice and noodles")
}

type American struct{}

func (a American) MyFavoriteFood() {
	fmt.Println("American like eating hamburger")
}

func NewEatFool(name string) EatFool {
	switch name {
	case "c":
		return Chinese{}
	case "a":
		return American{}
	}
	return nil
}
