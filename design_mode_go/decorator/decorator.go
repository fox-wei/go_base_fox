package decorator

import "fmt"

/*
 *装饰器模式(Decorator pattern)
 *在不改变原有对象的基础上，动态地给对象添加额外的职责；属于结构型设计模式

  ?给对象添加新行为，最简单的方式是继承；继承的弊端：1.静态的；2.子类只能有一个父类，添加太多功能，导致类剧增
*/

//^Component
type PS5 interface {
	StartGPUEngine()
	GetPrice() int64
}

//~ConreteComponent
type PS5WithCD struct{}

func (p PS5WithCD) StartGPUEngine() {
	fmt.Println("start engine")
}

func (p PS5WithCD) GetPrice() int64 {
	return 1500
}

type PS5WithDigital struct{}

func (p PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal you engine")
}

func (p PS5WithDigital) GetPrice() int64 {
	return 3600
}

//!Decorator
//*plus decorator
type PS5MachinePlus struct {
	ps5Machine PS5
}

func (p *PS5MachinePlus) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("start plus plugin")
}

func (p PS5MachinePlus) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 500
}

func Client() {
	plusMachine := PS5MachinePlus{}
	plusMachine.SetPS5Machine(PS5WithCD{})
	plusMachine.StartGPUEngine()
	prince := plusMachine.GetPrice()
	fmt.Printf("PS5 CD plus, price:%d\n", prince)
}
