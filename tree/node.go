package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	// 工厂方法
	return &Node{Value: value} // 返回局部变量地址
}

func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored")
		return
	}
	node.Value = value
}
