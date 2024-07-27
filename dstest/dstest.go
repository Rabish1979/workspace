package main

import (
	"fmt"

	L "github.com/Rabish1979/workspace/linklist"
)

func main() {
	var myPlaylist L.LinkList
	myPlaylist2 := L.createLinkList()
	myPlaylist2.insertAtFront(20)
	myPlaylist2.insertAtFront(10)
	myPlaylist2.print()

	fmt.Printf("myPlaylist2: %v\n", myPlaylist)
	fmt.Println("hello ds")
}
