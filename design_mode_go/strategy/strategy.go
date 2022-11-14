package strategy

import "fmt"

//*策略模式实现-支付方式模拟

//&抽象策略
type PayBehavior interface {
	OrderPay(*PayCtx)
}

//*具体策略-支付方式
type WxPay struct {
}

func (w *WxPay) OrderPay(p *PayCtx) {
	fmt.Printf("Wx支付加工请求，%v\n", p.payParams)
	fmt.Println("正在使用Wx支付")
}

type ThirdPay struct {
}

func (td *ThirdPay) OrderPay(p *PayCtx) {
	fmt.Printf("ThirdPay支付加工请求，%v\n", p.payParams)
	fmt.Println("正在使用ThirdPay支付")
}

//&上下文
type PayCtx struct {
	//*提供支付接口
	payBehavior PayBehavior

	//*支付参数
	payParams map[string]interface{}
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	//^Mock 数据
	params := map[string]interface{}{
		"appId": "333434",
		"mchId": 12345,
	}

	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}
