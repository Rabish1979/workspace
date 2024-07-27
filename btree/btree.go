package btree

import "fmt"

//import necessary packages fmt and strings

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int32
}

type BTree struct {
	head *TreeNode
}

func CreateBTree() *BTree {
	return &BTree{}
}

func (root *BTree) Insert(data int32) {
	if root.head == nil {
		newNode := &TreeNode{data: data, left: nil, right: nil}
		root.head = newNode
		return
	}

	var temp *TreeNode = root.head
	for temp != nil {
		if data < temp.data {
			temp = temp.left
		}

		if data > temp.data {
			temp = temp.right
		}
	}

	temp = &TreeNode{data: data, left: nil, right: nil}
}

func (root *BTree) InOrder() {
	root.inOrder(root.head)
}

func (root *BTree) inOrder(node *TreeNode) {
	var current *TreeNode = root.head

	if current.left != nil {
		root.inOrder(current.left)
	}

	fmt.Println("data =>", current.data)

	if current.right != nil {
		root.inOrder(current.right)
	}

	fmt.Println()
}
