package main

import "fmt"

type Node struct {
    Data int
    Prev *Node
    Next *Node
}

type CircularDoublyLinkedList struct {
    Head *Node
}

// Вставка в конец списка
func (c *CircularDoublyLinkedList) InsertAtEnd(data int) {
    newNode := &Node{Data: data}
    if c.Head == nil {
        c.Head = newNode
        c.Head.Prev = c.Head
        c.Head.Next = c.Head
        return
    }
    last := c.Head.Prev
    last.Next = newNode
    newNode.Prev = last
    newNode.Next = c.Head
    c.Head.Prev = newNode
}

// Вставка в начало списка
func (c *CircularDoublyLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{Data: data}
    if c.Head == nil {
        c.Head = newNode
        c.Head.Prev = c.Head
        c.Head.Next = c.Head
        return
    }
    last := c.Head.Prev
    newNode.Next = c.Head
    newNode.Prev = last
    c.Head.Prev = newNode
    last.Next = newNode
    c.Head = newNode
}

// Вставка после узла
func (c *CircularDoublyLinkedList) InsertAfter(node *Node, data int) {
    if node == nil {
        return
    }
    newNode := &Node{Data: data}
    nextNode := node.Next
    node.Next = newNode
    newNode.Prev = node
    newNode.Next = nextNode
    nextNode.Prev = newNode
}

// Удаление головы списка
func (c *CircularDoublyLinkedList) DeleteHead() {
    if c.Head == nil {
        return
    }
    if c.Head.Next == c.Head {
        c.Head = nil
        return
    }
    last := c.Head.Prev
    c.Head = c.Head.Next
    c.Head.Prev = last
    last.Next = c.Head
}

// Удаление узла по значению
func (c *CircularDoublyLinkedList) DeleteNodeByValue(data int) {
    if c.Head == nil {
        return
    }
    if c.Head.Data == data {
        c.DeleteHead()
        return
    }
    current := c.Head
    for {
        if current.Data == data {
            prevNode := current.Prev
            nextNode := current.Next
            prevNode.Next = nextNode
            nextNode.Prev = prevNode
            return
        }
        current = current.Next
        if current == c.Head {
            break
        }
    }
}

// Поиск узла по значению
func (c *CircularDoublyLinkedList) Search(data int) *Node {
    if c.Head == nil {
        return nil
    }
    current := c.Head
    for {
        if current.Data == data {
            return current
        }
        current = current.Next
        if current == c.Head {
            break
        }
    }
    return nil
}

// Итерация по списку (вперёд)
func (c *CircularDoublyLinkedList) IterationForward() {
    if c.Head == nil {
        fmt.Println("List is empty")
        return
    }
    current := c.Head
    for {
        fmt.Printf("%d ", current.Data)
        current = current.Next
        if current == c.Head {
            break
        }
    }
    fmt.Println()
}

// Итерация по списку (назад)
func (c *CircularDoublyLinkedList) IterationBackward() {
    if c.Head == nil {
        fmt.Println("List is empty")
        return
    }
    current := c.Head.Prev
    for {
        fmt.Printf("%d ", current.Data)
        current = current.Prev
        if current == c.Head.Prev {
            break
        }
    }
    fmt.Println()
}

func main() {
    list := &CircularDoublyLinkedList{}

    fmt.Println("---------------------------")
    fmt.Println("InsertAtEnd: 10, 20, 30")
    list.InsertAtEnd(10)
    list.InsertAtEnd(20)
    list.InsertAtEnd(30)
    list.IterationForward()

    fmt.Println("---------------------------")
    fmt.Println("InsertAtBeginning: 5")
    list.InsertAtBeginning(5)
    list.IterationForward()

    fmt.Println("---------------------------")
    fmt.Println("InsertAfter node with value 10: 15")
    node := list.Search(10)
    if node != nil {
        list.InsertAfter(node, 15)
    }
    list.IterationForward()

    fmt.Println("---------------------------")
    fmt.Println("DeleteNodeByValue: 20")
    list.DeleteNodeByValue(20)
    list.IterationForward()

    fmt.Println("---------------------------")
    fmt.Println("DeleteHead")
    list.DeleteHead()
    list.IterationForward()

    fmt.Println("---------------------------")
    fmt.Println("Search for 15:", list.Search(15) != nil)
    fmt.Println("Search for 99:", list.Search(99) != nil)

    fmt.Println("---------------------------")
    fmt.Println("IterationBackward:")
    list.IterationBackward()
}
