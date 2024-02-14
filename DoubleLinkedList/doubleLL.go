package main

import "fmt"

//DoubleLinkedList
//전 주소와, 다음 이어지는 주소를 동시에 가지는 리스트 형식이다.

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) removeNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}
	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}

}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.removeNode(list.root.next)

	list.PrintNodes()

	list.removeNode(list.root)

	list.PrintNodes()

	list.removeNode(list.tail)

	list.PrintNodes()

	fmt.Printf("tail: %d", list.tail.val)
}
