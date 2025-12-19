package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

// O(1)
func (l *LinkedList) InsertHead(data int) {
	newNode := &Node{
		Data: data,
		Next: l.Head,
	}

	l.Head = newNode
}

// O(n)
func (l *LinkedList) InsertTail(data int) {
	newNode := &Node{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// O(n)
func (l *LinkedList) InsertIndex(data int, index int) {
	if index == 0 {
		l.InsertHead(data)
		return
	}

	newNode := &Node{
		Data: data,
	}

	current := l.Head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current != nil {
		newNode.Next = current.Next
		current.Next = newNode
	}
}

// O(1)
func (l *LinkedList) InsertNode(position *Node, data int) {

	if position == nil {
		return
	}

	newNode := &Node{
		Data: data,
		Next: position.Next,
	}

	position.Next = newNode
}

// O(n)
func (l *LinkedList) InsertValue(target, data int) {

	current := l.Head

	for current != nil {
		if current.Data == target {
			newNode := &Node{
				Data: data,
				Next: current.Next,
			}

			current.Next = newNode
			return
		}
		current = current.Next
	}
}

// O(1)
func (l *LinkedList) DeleteHead() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

// O(n)
func (l *LinkedList) DeleteTail() {
	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		return
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
}

// O(n)
func (l *LinkedList) DeleteIndex(index int) {

	if index == 0 {
		l.DeleteHead()
		return
	}

	current := l.Head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}
}

// O(n)
func (l *LinkedList) DeleteNode(node *Node) {
	if node == nil || l.Head == nil {
		return
	}

	if l.Head == node {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current != nil && current.Next != node {
		current = current.Next
	}

	if current != nil {
		current.Next = current.Next.Next
	}
}

// O(n)
func (l *LinkedList) DeleteValue(val int) {

	if l.Head == nil {
		return
	}

	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// O(n)
func (l *LinkedList) Search(data int) bool {
	current := l.Head

	for current != nil {
		if current.Data == data {
			return true
		}

		current = current.Next
	}

	return false
}

// O(n)
func (l *LinkedList) Iteration() {
	current := l.Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
    list := LinkedList{}

    fmt.Println("InsertHead: 10, 20, 30")
    list.InsertHead(10)
    list.InsertHead(20)
    list.InsertHead(30)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("InsertTail: 55")
    list.InsertTail(55)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("InsertIndex: 77 at index 2")
    list.InsertIndex(77, 2)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("InsertNode: 88 after head")
    list.InsertNode(list.Head, 88)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("InsertValue: 99 after 77")
    list.InsertValue(77, 99)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("DeleteHead")
    list.DeleteHead()
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("DeleteTail")
    list.DeleteTail()
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("DeleteIndex: 2")
    list.DeleteIndex(2)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("DeleteNode: 99")
    current := list.Head
    for current != nil {
        if current.Data == 99 {
            break
        }
        current = current.Next
    }
    list.DeleteNode(current)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("DeleteValue: 88")
    list.DeleteValue(88)
    list.Iteration()
    fmt.Println("-------------------------------------------------")

    fmt.Println("Search: 77")
    fmt.Println(list.Search(77))
    fmt.Println("-------------------------------------------------")

    fmt.Println("Final iteration:")
    list.Iteration()
}
