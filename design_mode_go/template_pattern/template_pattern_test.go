package templatepattern

import (
	"testing"
)

func TestBankBusinessExector_ExecuteBusiness(t *testing.T) {
	deposit := DepositBusinessHandler{userVip: false}
	hanlder := NewBankBusinessExecutor(&deposit)
	hanlder.ExecuteBusiness()
}
