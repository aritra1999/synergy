package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) insert(key int) {
	if bst.root == nil {
		bst.root = &Node{key: key}
	} else {
		bst.root.insert(key)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func (node *Node) insert(key int) {
	if key <= node.key {
		if node.left == nil {
			node.left = &Node{key: key}
		} else {
			node.left.insert(key)
		}
	} else {
		if node.right == nil {
			node.right = &Node{key: key}
		} else {
			node.right.insert(key)
		}
	}
}

func main() {
	var bst BST
	bst.insert(10)
	bst.insert(5)
	bst.insert(15)
	bst.insert(8)
	bst.insert(3)
	bst.insert(12)

	printPreOrder(bst.root)

}
