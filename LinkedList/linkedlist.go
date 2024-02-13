package main

import "fmt"

//Linkedlist 예제

type Node struct {
	next *Node
	val  int
}

func main() {
	//처음 시작 노드
	var root *Node
	var tail *Node
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	root, tail = removeNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = removeNode(root, root, tail)

	PrintNodes(root)

	root, tail = removeNode(tail, root, tail)

	PrintNodes(root)
	fmt.Printf("tail: %d",tail.val)


}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func removeNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
