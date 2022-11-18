package templatepattern

import (
	"fmt"
	"math/rand"
	"time"
)

type BankBusinessHandler interface {
	//*取号
	TakeNum()

	//*等待
	waitHead()

	//&处理业务
	HandleBusiness()

	//*评价
	Comment()

	//&VIP
	CheckVip() bool
}

//*模板方法
type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

func (b *BankBusinessExecutor) ExecuteBusiness() {
	//*使用于与客户端单次交互
	//*如果需要多次交互才能完成流程，每次操作写入对应模板方法中
	b.handler.TakeNum()
	if !b.handler.CheckVip() {
		b.handler.waitHead()
	}
	b.handler.HandleBusiness()
	b.handler.Comment()
}

//*取款业务
type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("xxx存款xx元")
}

func (dbh *DepositBusinessHandler) CheckVip() bool {
	return dbh.userVip
}

type DefaultBusinessHandler struct{}

func (*DefaultBusinessHandler) TakeNum() {
	fmt.Printf("排队码：%d，注意排队\n", rand.Intn(100))
}

func (*DefaultBusinessHandler) waitHead() {
	fmt.Println("排队等号中...+")
	time.Sleep(3 * time.Second)
	fmt.Println("请xx到xx窗口...")
}

func (*DefaultBusinessHandler) Comment() {
	fmt.Println("xxx请对服务做出评价...")
}

func (*DefaultBusinessHandler) CheckVip() bool {
	//*具体类实现
	return false
}

func NewBankBusinessExecutor(business BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{business}
}
