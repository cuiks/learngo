package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	// 工厂方法
	return &treeNode{value: value} // 返回局部变量地址
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)
	root.right.left.print()

	fmt.Println("root: ")
	root.print()
	root.setValue(100)
	root.print()

	fmt.Println("pRoot: ")
	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var nRoot *treeNode
	nRoot.setValue(200)
	nRoot = &root
	nRoot.setValue(300)
	nRoot.print()

	root.traverse()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
}
