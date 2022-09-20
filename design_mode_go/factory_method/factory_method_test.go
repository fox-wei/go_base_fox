package factory_method

import (
	"fmt"
	"testing"
)

func TestMultiOperator_ComputeResult(t *testing.T) {
	var factory OperatorFactory
	var mathOp MathOperator
	//*加法操作
	factory = &PlusFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandA(23)
	mathOp.SetOperandB(24)
	fmt.Println(mathOp.ComputeResult())

	//*乘法
	factory = &MultiFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandA(4)
	mathOp.SetOperandB(7)
	fmt.Println(mathOp.ComputeResult())
}
