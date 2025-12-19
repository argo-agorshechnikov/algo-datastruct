package main

import "fmt"

type Node struct {
    Data int
    Next *Node
}

type CircularLinkedList struct {
    Head *Node
}

// Вставка в конец списка
func (c *CircularLinkedList) InsertAtEnd(data int) {
    newNode := &Node{Data: data}
    if c.Head == nil {
        c.Head = newNode
        c.Head.Next = c.Head
        return
    }
    current := c.Head
    for current.Next != c.Head {
        current = current.Next
    }
    current.Next = newNode
    newNode.Next = c.Head
}

// Вставка в начало списка
func (c *CircularLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{Data: data}
    if c.Head == nil {
        c.Head = newNode
        c.Head.Next = c.Head
        return
    }
    current := c.Head
    for current.Next != c.Head {
        current = current.Next
    }
    current.Next = newNode
    newNode.Next = c.Head
    c.Head = newNode
}

// Вставка после узла
func (c *CircularLinkedList) InsertAfter(node *Node, data int) {
    if node == nil {
        return
    }
    newNode := &Node{Data: data, Next: node.Next}
    node.Next = newNode
}

// Удаление головы списка
func (c *CircularLinkedList) DeleteHead() {
    if c.Head == nil {
        return
    }
    if c.Head.Next == c.Head {
        c.Head = nil
        return
    }
    current := c.Head
    for current.Next != c.Head {
        current = current.Next
    }
    current.Next = c.Head.Next
    c.Head = c.Head.Next
}

// Удаление узла по значению
func (c *CircularLinkedList) DeleteNodeByValue(data int) {
    if c.Head == nil {
        return
    }
    if c.Head.Data == data {
        c.DeleteHead()
        return
    }
    current := c.Head
    for current.Next != c.Head {
        if current.Next.Data == data {
            current.Next = current.Next.Next
            return
        }
        current = current.Next
    }
}

// Удаление узла по указателю
func (c *CircularLinkedList) DeleteNode(node *Node) {
    if node == nil || c.Head == nil {
        return
    }
    if c.Head == node {
        c.DeleteHead()
        return
    }
    current := c.Head
    for current.Next != c.Head {
        if current.Next == node {
            current.Next = current.Next.Next
            return
        }
        current = current.Next
    }
}

// Поиск узла по значению
func (c *CircularLinkedList) Search(data int) *Node {
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

// Итерация по списку
func (c *CircularLinkedList) Iteration() {
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

func main() {
    list := &CircularLinkedList{}

    fmt.Println("---------------------------")
    fmt.Println("InsertAtEnd: 10, 20, 30")
    list.InsertAtEnd(10)
    list.InsertAtEnd(20)
    list.InsertAtEnd(30)
    list.Iteration()

    fmt.Println("---------------------------")
    fmt.Println("InsertAtBeginning: 5")
    list.InsertAtBeginning(5)
    list.Iteration()

    fmt.Println("---------------------------")
    fmt.Println("InsertAfter node with value 10: 15")
    node := list.Search(10)
    if node != nil {
        list.InsertAfter(node, 15)
    }
    list.Iteration()

    fmt.Println("---------------------------")
    fmt.Println("DeleteNodeByValue: 20")
    list.DeleteNodeByValue(20)
    list.Iteration()

    fmt.Println("---------------------------")
    fmt.Println("DeleteHead")
    list.DeleteHead()
    list.Iteration()

    fmt.Println("---------------------------")
    fmt.Println("Search for 15:", list.Search(15) != nil)
    fmt.Println("Search for 99:", list.Search(99) != nil)

    fmt.Println("---------------------------")
    fmt.Println("DeleteNode: node with value 15")
    node = list.Search(15)
    if node != nil {
        list.DeleteNode(node)
    }
    list.Iteration()
}
