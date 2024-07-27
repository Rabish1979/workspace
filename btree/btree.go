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

func (btree *BTree) Insert(data int32) {
	if btree.head == nil {
		btree.head = &TreeNode{data: data, left: nil, right: nil}
	} else {
		btree.head.insert(data)
	}
}

func (root *TreeNode) insert(data int32) {
	if data <= root.data {
		if root.left == nil {
			root.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			root.left.insert(data)
		}

	} else if data > root.data {
		if root.right == nil {
			root.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			root.right.insert(data)
		}
	}
}

func (root *BTree) InOrder() {
	root.head.inOrder()
}

func (root *TreeNode) inOrder() {
	if root == nil {
		return
	}

	if root.left != nil {
		root.left.inOrder()
	}

	fmt.Println("data =>", root.data)

	if root.right != nil {
		root.right.inOrder()
	}

	fmt.Println()
}
