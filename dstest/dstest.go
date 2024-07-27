package main

import (
	"fmt"

	T "github.com/Rabish1979/workspace/btree"
	L "github.com/Rabish1979/workspace/linklist"
)

func main() {
	var myPlaylist L.LinkList
	myPlaylist2 := L.CreateLinkList()
	myPlaylist2.InsertAtFront(20)
	myPlaylist2.InsertAtFront(10)
	myPlaylist2.Print()

	fmt.Printf("myPlaylist2: %v\n", myPlaylist)
	fmt.Println("hello ds")

	myTree := T.CreateBTree()
	myTree.Insert(23)
	myTree.Insert(20)
	myTree.Insert(19)
	myTree.Insert(25)
	myTree.Insert(30)
	myTree.InOrder()
}
