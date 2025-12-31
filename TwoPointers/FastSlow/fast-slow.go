package main

import "fmt"

// Detect cycle in Linked List

type Node struct {
	Val int
	Next *Node
}

func hasCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	cycleList := &Node{3, nil}
	cycleList.Next = &Node{1, nil}
    cycleList.Next.Next = &Node{2, nil}
    cycleList.Next.Next.Next = &Node{4, nil}
    cycleList.Next.Next.Next.Next = cycleList.Next.Next

	fmt.Println("cycleList: ", hasCycle(cycleList))

	linearList := &Node{1, nil}
    linearList.Next = &Node{2, nil}
    linearList.Next.Next = &Node{3, nil}
	fmt.Println("linearList:", hasCycle(linearList))
}
