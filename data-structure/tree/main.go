package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &Node{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.Root.insertNode(newNode)
	}
}

func (n *Node) insertNode(newNode *Node) {
	if newNode.Value < n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.insertNode(newNode)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.insertNode(newNode)
		}
	}
}
func (bt *BinarySearchTree) Search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.Value > value {
		fmt.Println(root, bt.Root)
		return bt.Search(root.Left, value)
	} else if root.Value < value {
		return bt.Search(root.Right, value)
	}
	return true
}
func DeleteNode(n *Node, value int) *Node {
	if n == nil {
		return nil
	}
	if n.Value > value {
		n.Left = DeleteNode(n.Left, value)
	} else if n.Value < value {
		n.Right = DeleteNode(n.Right, value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		MinRight := FindMin(n.Right)
		n.Value = MinRight.Value
		n.Right = DeleteNode(n.Right, MinRight.Value)
	}
	return n
}
func FindMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
func Height(node *Node) int {
	if node == nil {
		return 0
	}
	left := Height(node.Left)
	right := Height(node.Right)
	return max(left, right) + 1
}
func balanced(node *Node) bool {
	if node == nil {
		return true
	}
	left := Height(node.Left)
	right := Height(node.Right)
	if abs(left-right) > 1 {
		return false
	}
	return true
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	bst := &BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(15)
	bst.Insert(12)
	bst.Insert(18)
	bst.Insert(112)
	bst.Insert(1434)

	fmt.Println("height", Height(bst.Root))
	DeleteNode(bst.Root, 3)
	Postorder(bst.Root)
	if balanced(bst.Root) {
		fmt.Println("balanced")
	} else {
		fmt.Println("unbalanced")
	}
	// fmt.Println("Search 7:", bst.Search(bst.Root, 5))
	// fmt.Println("Search 20:", bst.Search(20))
}
func InoderTraversal(root *Node) {
	if root != nil {
		InoderTraversal(root.Left)
		fmt.Println(root.Value)
		InoderTraversal(root.Right)
	}
}
func Preorder(node *Node) {
	if node != nil {
		fmt.Println(node.Value)
		Preorder(node.Left)
		Preorder(node.Right)
	}
}
func Postorder(node *Node) {
	if node != nil {
		Preorder(node.Left)
		Preorder(node.Right)
		fmt.Println(node.Value)
	}
}
