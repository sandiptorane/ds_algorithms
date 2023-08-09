package tree

import (
	"fmt"
)

// BinaryNode holds value of info of left, right node
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

// BinaryTree holds root node
type BinaryTree struct {
	root *BinaryNode
}

// Insert : add data with node to binary tree
func (b *BinaryTree) Insert(data int) *BinaryTree {
	// if root not nil add data to next nodes
	// else add to root
	if b.root != nil {
		b.root.Insert(data)
		return b
	}

	b.root = &BinaryNode{data: data}

	return b
}

// Insert data to either left or right based on value
func (n *BinaryNode) Insert(data int) {
	// if data in less than node data insert to left else insert to right
	if data <= n.data {
		if n.left != nil {
			n.left.Insert(data)
			return
		}

		n.left = &BinaryNode{data: data}
		fmt.Printf("%v inserted to left of %v\n", data, n.data)
		return
	}

	if n.right != nil {
		n.right.Insert(data)
		return
	}

	n.right = &BinaryNode{data: data}
	fmt.Printf("%v inserted to right of %v\n", data, n.data)
}

// Print binary tree in order
func (b *BinaryTree) Print() {
	b.root.Print()
}

func (b *BinaryNode) Print() {
	if b == nil {
		return
	}

	b.left.Print()
	fmt.Printf("%+v ", b.data)
	b.right.Print()
}
