package factory_method

/*
*工厂方法模式：定义一个创建对象的接口，但由实现这个接口的工厂
*类来决定实例化哪一个产品类。
 */

//&以工厂生产各种类型计算器为例

//&抽象工厂
type OperatorFactory interface {
	Create() MathOperator
}

//*抽象产品
type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	ComputeResult() int
}

//?BaseOperator是所有operator的基类
type BaseOperator struct {
	operandA, operandB int
}

func (b *BaseOperator) SetOperandA(operand int) {
	b.operandA = operand
}

func (b *BaseOperator) SetOperandB(operand int) {
	b.operandB = operand
}

//!具体工厂-加法运算
type PlusFactory struct{}

func (pf *PlusFactory) Create() MathOperator {
	return &PlusOPerator{
		BaseOperator: &BaseOperator{},
	}
}

type MultiFactory struct{}

func (mf *MultiFactory) Create() MathOperator {
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

//!具体产品
//*加法运算
type PlusOPerator struct {
	*BaseOperator
}

func (p *PlusOPerator) ComputeResult() int {
	return p.operandA + p.operandB
}

//*乘法运算
type MultiOperator struct {
	*BaseOperator
}

func (m *MultiOperator) ComputeResult() int {
	return m.operandA * m.operandB
}
