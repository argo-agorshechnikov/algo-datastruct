package main

import "fmt"


type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{Value: val}
	}

	if val < n.Value {
		n.Left = (*Node)(n.Left).Insert(val)
	} else if val > n.Value {
		n.Right = (*Node)(n.Right).Insert(val)
	}

	return n
}


func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if val == n.Value {
		return true
	}

	if val < n.Value {
		return (*Node)(n.Left).Search(val)
	}

	return (*Node)(n.Right).Search(val)
}

func (n *Node) Min() *Node {
	if n == nil || n.Left == nil {
		return n
	}

	return (*Node)(n.Left).Min()
}

func (n *Node) Delete(val int) *Node {
	if n == nil {
		return nil
	}

	if val < n.Value {
		n.Left = (*Node)(n.Left).Delete(val)
	} else if val > n.Value {
		n.Right = (*Node)(n.Right).Delete(val)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		minNode := (*Node)(n.Right).Min()
		n.Value = minNode.Value
		n.Right = (*Node)(n.Right).Delete(minNode.Value)
	}

	return n
	
}


func (n *Node) InOrder() {
	if n != nil {
		(*Node)(n.Left).InOrder()
		fmt.Println("value: ", n.Value)
		(*Node)(n.Right).InOrder()
	}
}

func main() {
	root := (*Node)(nil).Insert(50)
    root.Insert(30)
    root.Insert(20)
    root.Insert(40)
    root.Insert(70)
    root.Insert(60)
    root.Insert(80)

	fmt.Print("InOrder: ")
	root.InOrder()
	fmt.Println("-------------------------------")

	fmt.Println("Search 40: ", root.Search(40))

	fmt.Println("-------------------------------")
	fmt.Println("Delete 40")
	root.Delete(40)
	root.InOrder()
}