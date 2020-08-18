package main

import "fmt"

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2
	head.next = node1
	var node2 = new(Node)
	node2.data = 3
	node1.next = node2
	ShowList(head)
}

type Node struct {
	data int
	next *Node
}

// 遍历
func ShowList(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}