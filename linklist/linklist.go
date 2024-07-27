package linklist

import "fmt"

//import necessary packages fmt and strings

type Node struct {
	next *Node
	data int32
}

type LinkList struct {
	head *Node
}

func createLinkList() *LinkList {
	return &LinkList{}
}

func (list *LinkList) insertAtFront(data int32) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println()
}
