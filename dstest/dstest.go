package main

import (
	"fmt"
)

func main() {
	myPlaylist := createLinkList()
	myPlaylist.insertAtFront(20)
	myPlaylist.insertAtFront(30)
	myPlaylist.insertAtFront(40)
	myPlaylist.Print()
	fmt.Println("hello ds")
}
