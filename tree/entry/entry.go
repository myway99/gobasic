package main

import (
	"fmt"

	"../../tree"
	"golang.org/x/tools/container/intsets"
)

//后序遍历
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()

	myNode.node.Print()

}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	//数节点
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max Node value:", maxNode)

	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	//fmt.Println()
	//
	//testSparse()

	//
	//root.right.left.setValue(4)
	//root.right.left.print()
	//fmt.Println()
	//
	//root.print()
	//root.setValue(100)
	//
	//var pRoot *treeNode
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)

}
