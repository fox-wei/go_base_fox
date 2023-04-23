package composite

import (
	"fmt"
	"strings"
)

//*组合模式(Composite Pattern)：部分-整体（Part-Whole）模式，它的宗旨是
//*通过对个对象（叶子节点）和组合对象（树枝节点）用相同接口进行表示，使得对
//*单个对象和组合对象使用具有一致性，属于结构型设计模式。
//?应用场景：对象能组织成树形结构
/*
模式角色
 ~ 组件（Componment）：组件是一个接口，描述树中对象的操作
 ~ 叶子结点（Leaf）：单个节点对象，基本结构
 ~ 组合对象（Composite）：包含叶子节点或其他组合对象等符合项目的对象
 ~ 客户端（Clinet）：组件与接口交互
*/

//?组织结构-组件
type Organization interface {
	Display() //^组织结构
	Duty()    //&职责
}

//?组合对象
type CompoisteOrganization struct {
	orgName string
	depth   int
	list    []Organization
}

func NewCompositeOrganization(name string, depth int) *CompoisteOrganization {
	return &CompoisteOrganization{name, depth, []Organization{}}
}

func (c *CompoisteOrganization) Add(org Organization) {
	if c == nil {
		return
	}
	c.list = append(c.list, org)
	// c.depth += 1
}

func (c *CompoisteOrganization) Remove(org Organization) int {
	if c == nil {
		return 0
	}

	for i, v := range c.list {
		if v == org {
			c.list = append(c.list[:i], c.list[i+1:]...)
			// c.depth -= 1
			return 1
		}
	}

	return 0
}

func (c *CompoisteOrganization) Display() {
	if c == nil {
		return
	}

	fmt.Println(strings.Repeat("-", c.depth), c.orgName)
	for _, v := range c.list {
		v.Display()
	}
}

func (c *CompoisteOrganization) Duty() {
	if c == nil {
		return
	}

	for _, v := range c.list {
		v.Duty()
	}
}

//!职能部门实现-叶子
//&人力资源
type HROrg struct {
	orgName string
	depth   int
}

func (h *HROrg) Display() {
	if h == nil {
		return
	}

	fmt.Println(strings.Repeat("-", h.depth*2), " ", h.orgName)
}

func (h *HROrg) Duty() {
	if h == nil {
		return
	}
	fmt.Println(h.orgName, " 员工招聘管理部门")
}

//&财务部门
type FinanceOrg struct {
	orgName string
	depth   int
}

func (f *FinanceOrg) Display() {
	fmt.Println(strings.Repeat("-", f.depth*2), " ", f.orgName)
}

func (f *FinanceOrg) Duty() {
	fmt.Println(f.orgName, " 财务部门管理")
}
