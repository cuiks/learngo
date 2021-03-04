package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets"
	"learngo/tree"
)

// 通过组合的方式实现扩展方法
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)

	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	//root.Right.Left.Print()

	//fmt.Println("root: ")
	//root.Print()
	//root.SetValue(100)
	//root.Print()
	//
	//fmt.Println("pRoot: ")
	//pRoot := &root
	//pRoot.Print()
	//pRoot.SetValue(200)
	//pRoot.Print()
	//
	//var nRoot *tree.Node
	//nRoot.SetValue(200)
	//nRoot = &root
	//nRoot.SetValue(300)
	//nRoot.Print()
	//
	root.Traverse()
	//
	//fmt.Println("###")
	//myNode := myTreeNode{&root}
	//myNode.postOrder()
	//fmt.Println("###")

	//nodes := []tree.TreeNode{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//testSparse()

	//fmt.Println("####")
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count: ", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
