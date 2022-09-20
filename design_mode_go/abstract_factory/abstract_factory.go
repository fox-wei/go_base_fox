package abstract_factory

import "fmt"

//*抽象工厂模式：用于创建一系列相关的或者相互依赖的对象
//?工厂方法模式针对的是一个产品的等级结构
//^抽象工厂模式则需要面对多个产品等级结构

//?抽象工厂
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

//*抽象产品
type ITelevision interface {
	Wathch()
}

type IAirConditioner interface {
	SetTemperature(int)
}

//*具体工厂和具体产品族

//&华为工厂
type HuaweiFactory struct{}

func (hf *HuaweiFactory) CreateTelevision() ITelevision {
	return &HuaweiTv{}
}

func (hf *HuaweiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaweiAirConditioner{}
}

type HuaweiTv struct{}

func (hw *HuaweiTv) Wathch() {
	fmt.Println("watch huawei TV")
}

type HuaweiAirConditioner struct{}

func (ha *HuaweiAirConditioner) SetTemperature(temperature int) {
	fmt.Printf("The hua wei aircondition's temperature is %d\n", temperature)
}

//*Mi工厂
type MiFactory struct{}

func (mf *MiFactory) CreateTelevision() ITelevision {
	return &MiTelevision{}
}

func (mf *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}

type MiTelevision struct{}

func (mt *MiTelevision) Wathch() {
	fmt.Println("wath Mi Tv")
}

type MiAirConditioner struct{}

func (ma *MiAirConditioner) SetTemperature(temperature int) {
	fmt.Printf("The Mi aircondition's temperature is %d\n", temperature)
}
