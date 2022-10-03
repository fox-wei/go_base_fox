package prototype

import (
	"fmt"
	"testing"
)

func TestNewElement(t *testing.T) {
	//*创建公司职级关系

	//?职级关系-总监
	directorNode := NewElement("Director of Engineering")

	//?研发经理
	engManagerNode := NewElement("Engineering Manager")
	engManagerNode.AddChild(NewElement("Lead Sofwtare Engineer"))

	//*研发经理是总监的下级
	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(engManagerNode)

	//*办公室经理
	officeManageNode := NewElement("Office Manager")
	directorNode.AddChild(officeManageNode)

	fmt.Println("------")
	fmt.Println("# Company Hierarchy")
	fmt.Println(directorNode)

	fmt.Println("==从研发经理节点克隆出一颗新树==")
	fmt.Println("$ Team Hierarchy")
	fmt.Println(engManagerNode.Clone())

}
