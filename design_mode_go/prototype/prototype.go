package prototype

import (
	"bytes"
	"fmt"
)

//*原型模式:通过克隆已有对象的方式来创建对象，被拷贝对象称为原型对象
//&节省创建对象所花费的时间、资源，提供性能

//!实现文档树(DOM树)

//&DOM对象
type Node interface {
	//?信息描述
	String() string

	//?父节点
	Parent() Node

	SetParent(node Node)

	//*返回孩子节点
	Children() []Node

	AddChild(child Node)

	Clone() Node
}

type Element struct {
	text     string
	parent   Node
	children []Node
}

func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

//*implement the Node  interface
func (e *Element) Parent() Node {
	return e.parent
}

func (e *Element) SetParent(node Node) {
	e.parent = node
}

func (e *Element) Children() []Node {
	return e.children
}

func (e *Element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.children {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}

func (e *Element) AddChild(child Node) {
	copy := child.Clone()
	copy.SetParent(e)
	e.children = append(e.children, copy)
}

func (e *Element) Clone() Node {
	copy := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}

	for _, child := range e.children {
		copy.AddChild(child)
	}

	return copy
}
