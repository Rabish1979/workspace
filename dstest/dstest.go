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

	TestHashMap()
}

func TestHashMap() {
	hashmap := make(map[string]int)

	// Add some items to the map
	hashmap["pencil"] = 10
	hashmap["pen"] = 20
	hashmap["scale"] = 15

	// Print the entire map
	fmt.Println("The hashmap created above is: ")
	fmt.Println(hashmap)

	// Access a specific item by its key
	fmt.Println("The value of specific element from hashmap is:")
	fmt.Println(hashmap["pencil"])

	// Update the value of an existing item
	hashmap["scale"] = 20

	// Delete an item from the map
	delete(hashmap, "pencil")

	// Print the updated map
	fmt.Println("The hashmap after deleting an item from it is:")
	fmt.Println(hashmap)
}
