package factory

import "fmt"

/*
*简单工厂模式：创建对象实例
 */

//?定义接口，吃
//?抽象产品
type EatFool interface {
	MyFavoriteFood()
}

//*具体产品
type Chinese struct{}

func (c Chinese) MyFavoriteFood() {
	fmt.Println("Chinese like eating the rice and noodles")
}

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
