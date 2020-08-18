package main

import "fmt"

func main() {
	// 创建根节点
	root := NewNode(nil, nil)
	var it Initer
	it = root
	it.SetData("root")

	// 创建左子树
	a := NewNode(nil, nil)
	a.SetData("left")
	al := NewNode(nil, nil)     // 左叶子节点
	al.SetData(100)
	ar := NewNode(nil, nil)     // 右叶子节点
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar

	// 创建右子树
	b := NewNode(nil, nil)
	b.SetData("right")
	root.Left = a
	root.Right = b

	root.PrintBT()

	fmt.Println()
	fmt.Println("The depths of the Btree is:", root.Depth())
	fmt.Println("The leaf counts of the Btree is:", root.LeafCount())

	root.PreOrder()       // 先序遍历
	fmt.Println()
	root.InOrder()        // 中序遍历
	fmt.Println()
	root.PostOrder()      // 后序遍历
}

type Node struct {
	Left *Node
	Data interface{}
	Right *Node
}

type Initer interface {
	SetData (data interface{})
}

type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

func (n *Node) SetData(data interface{}) {
	n.Data = data
}

func (n *Node) PrintBT() {
	PrintBT(n)
}

func (n *Node) Depth() int {
	return Depth(n)
}

func (n *Node) LeafCount() int {
	return LeafCount(n)
}

func (n *Node) PreOrder() {
	PreOrder(n)
}

func (n *Node) InOrder() {
	InOrder(n)
}

func (n *Node) PostOrder() {
	PostOrder(n)
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func PrintBT(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			PrintBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			PrintBT(n.Right)
			fmt.Printf(")")
		}
	}
}

func Depth(n *Node) int {
	var depleft, depright int
	if n == nil {
		return 0
	} else {
		depleft = Depth(n.Left)
		depright = Depth(n.Right)
		if depleft > depright {
			return depleft + 1
		} else {
			return depright + 1
		}
	}
}

func LeafCount(n *Node) int {
	if n == nil {
		return 0
	} else if (n.Left == nil) && (n.Right == nil) {
		return 1
	} else {
		return (LeafCount(n.Left) + LeafCount(n.Right))
	}
}

func PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func InOrder(n *Node) {
	if n != nil {
		PreOrder(n.Left)
		fmt.Printf("%v ", n.Data)
		PreOrder(n.Right)
	}
}

func PostOrder(n *Node) {
	PreOrder(n.Left)
	PreOrder(n.Right)
	fmt.Printf("%v ", n.Data)
}