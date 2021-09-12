package facade

import (
	"fmt"
	"math"
	"strconv"
)

type CarMode struct{}

func NewCarMode() *CarMode {
	return &CarMode{}
}

func (c *CarMode) SetMode() {
	fmt.Println("carMode---mode...")
}

type CarEngine struct{}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (C *CarEngine) SetEngine() {
	fmt.Println("CarEngine---engine...")
}

type CarBody struct{}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Println("CarBody---body...")
}

type CarFacade struct {
	mode   CarMode
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		mode:   CarMode{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.mode.SetMode()
	c.engine.SetEngine()
	c.body.SetBody()
}

func ChangeNumber(n int, number string) (res int) {
	switch n {
	case 2:

		for i, j := 0, len(number)-1; i < len(number); i++ {
			tem, _ := strconv.Atoi(string(number[j]))
			res += tem * baseOp(2, i)
			j--
		}
	case 8:
		for i, j := 0, len(number)-1; i < len(number)-1; i++ {
			tem, _ := strconv.Atoi(string(number[j]))
			res += tem * baseOp(8, i)
			j--
		}
	case 16:
		for i, j := 0, len(number)-1; i < len(number)-2; i++ {
			tem, _ := strconv.Atoi(string(number[j]))
			res += tem * baseOp(16, i)
			j--
		}
	}
	return
}

func baseOp(base int, v int) int {
	return int(math.Pow(float64(base), float64(v)))
}

//*转换成10进制
func DecBase(value int) int {
	s := strconv.Itoa(value)
	if s[0] == '0' {
		if s[1] == 'x' || s[1] == 'X' {
			return ChangeNumber(16, s)
		}
		return ChangeNumber(8, s)
	}
	return ChangeNumber(2, s)
}
