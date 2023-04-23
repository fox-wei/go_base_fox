package composite

import (
	"fmt"
	"testing"
)

func TestCompoiste(t *testing.T) {
	root := NewCompositeOrganization("北京总公司", 1)
	root.Add(&HROrg{orgName: "总公司人力资源部", depth: 2})
	root.Add(&FinanceOrg{orgName: "总公司人力资源部", depth: 2})

	comSh := NewCompositeOrganization("上海分公司", 2)
	comSh.Add(&HROrg{orgName: "上海分公司人力资源部", depth: 3})
	comSh.Add(&FinanceOrg{orgName: "上海分公司财务部", depth: 3})
	root.Add(comSh)

	compGd := NewCompositeOrganization("广东分公司", 2)
	compGd.Add(&HROrg{orgName: "广东分公司人力资源部", depth: 3})
	compGd.Add(&FinanceOrg{orgName: "广东分公司财务部", depth: 3})
	root.Add(compGd)

	fmt.Println("公司架构:")
	root.Display()

	fmt.Println("各部职责:")
	root.Duty()
}
