package strategy

import "testing"

func TestOrderPay(t *testing.T) {
	wxPay := WxPay{}
	px := NewPayCtx(&wxPay)
	px.Pay()

	//?切换支付方式
	px.setPayBehavior(&ThirdPay{})
	px.Pay()
}
