package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// O(n)
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{
			Value: value,
		}
	}

	if n.Left == nil {
		n.Left = &Node{Value: value}
		return n
	}

	if n.Right == nil {
		n.Right = &Node{Value: value}
		return n
	}

	n.Left = n.Left.Insert(value)
	return n
}

// O(n)
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if n.Value == value {
		return true
	}

	return n.Left.Search(value) || n.Right.Search(value)
}

// O(n)
func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if n.Left != nil && n.Left.Value == value {
		n.Left = nil
		return n
	}

	if n.Right != nil && n.Right.Value == value {
		n.Right = nil
		return n
	}

	n.Left = n.Left.Delete(value)
	n.Right = n.Right.Delete(value)
	return n
}

// O(n)
func (n *Node) Size() int {
	if n == nil {
		return 0
	}

	return 1 + n.Left.Size() + n.Right.Size()
}

// O(n)
func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// O(n)
func (n *Node) PreOrder() {
	if n != nil {
		fmt.Println("val: ", n.Value)
		n.Left.PreOrder()
		n.Right.PreOrder()
	}
}

// O(n)
func (n *Node) InOrder() {
	if n != nil {
		n.Left.InOrder()
		fmt.Println("val: ", n.Value)
		n.Right.InOrder()
	}
}

// O(n)
func (n *Node) PostOrder() {
	if n != nil {
		n.Left.PostOrder()
		n.Right.PostOrder()
		fmt.Println("val: ", n.Value)
	}
}

// O(n)
func (n *Node) LevelOrder() {

	if n == nil {
		return
	}

	queue := list.New()
	queue.PushBack(n)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		node := elem.Value.(*Node)
		fmt.Println("val: ", node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

func main() {

	root := (*Node)(nil).
		Insert(1).
		Insert(2).
		Insert(3).
		Insert(4).
		Insert(5)

	fmt.Println("Size: ", root.Size())
	fmt.Println("Height: ", root.Height())

	fmt.Println("----------------------------")
	fmt.Print("PreOrder: ")
	root.PreOrder()

	fmt.Println("----------------------------")
	fmt.Print("InOrder: ")
	root.InOrder()

	fmt.Println("----------------------------")
	fmt.Print("PostOrder: ")
	root.PostOrder()

	fmt.Println("----------------------------")
	fmt.Print("LevelOrder: ")
	root.LevelOrder()

	fmt.Println("----------------------------")
	fmt.Println("Search: ", root.Search(3))

	fmt.Println("----------------------------")
	root.Delete(4)
	fmt.Println("Delete: ", root.Size())
}
