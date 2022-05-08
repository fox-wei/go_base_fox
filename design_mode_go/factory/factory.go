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

//*产品
type American struct{}

func (a American) MyFavoriteFood() {
	fmt.Println("American like eating hamburger")
}

//^工厂
func NewEatFool(name string) EatFool {
	switch name {
	case "c":
		return Chinese{}
	case "a":
		return American{}
	}
	return nil
}
