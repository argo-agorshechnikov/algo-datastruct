package main

import "fmt"

type Node struct {
    Data int
    Prev *Node
    Next *Node
}

type DoublyLinkedList struct {
    Head *Node
    Tail *Node
}

// Вставка в конец списка
func (d *DoublyLinkedList) InsertAtEnd(data int) {
    newNode := &Node{Data: data}
    if d.Head == nil {
        d.Head = newNode
        d.Tail = newNode
        return
    }
    newNode.Prev = d.Tail
    d.Tail.Next = newNode
    d.Tail = newNode
}

// Вставка в начало списка
func (d *DoublyLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{Data: data}
    if d.Head == nil {
        d.Head = newNode
        d.Tail = newNode
        return
    }
    newNode.Next = d.Head
    d.Head.Prev = newNode
    d.Head = newNode
}

// Вставка после узла
func (d *DoublyLinkedList) InsertAfter(node *Node, data int) {
    if node == nil {
        return
    }
    newNode := &Node{Data: data}
    newNode.Prev = node
    newNode.Next = node.Next
    if node.Next != nil {
        node.Next.Prev = newNode
    } else {
        d.Tail = newNode
    }
    node.Next = newNode
}

// Удаление головы списка
func (d *DoublyLinkedList) DeleteHead() {
    if d.Head == nil {
        return
    }
    if d.Head == d.Tail {
        d.Head = nil
        d.Tail = nil
        return
    }
    d.Head = d.Head.Next
    d.Head.Prev = nil
}

// Удаление узла по значению
func (d *DoublyLinkedList) DeleteNodeByValue(data int) {
    if d.Head == nil {
        return
    }
    if d.Head.Data == data {
        d.DeleteHead()
        return
    }
    current := d.Head
    for current != nil {
        if current.Data == data {
            if current == d.Tail {
                d.Tail = current.Prev
                d.Tail.Next = nil
            } else {
                current.Prev.Next = current.Next
                current.Next.Prev = current.Prev
            }
            return
        }
        current = current.Next
    }
}

// Поиск узла по значению
func (d *DoublyLinkedList) Search(data int) *Node {
    current := d.Head
    for current != nil {
        if current.Data == data {
            return current
        }
        current = current.Next
    }
    return nil
}

// Итерация по списку (вперёд)
func (d *DoublyLinkedList) IterationForward() {
    if d.Head == nil {
        fmt.Println("List is empty")
        return
    }
    current := d.Head
    for current != nil {
        fmt.Printf("%d ", current.Data)
        current = current.Next
    }
    fmt.Println()
}

// Итерация по списку (назад)
func (d *DoublyLinkedList) IterationBackward() {
    if d.Tail == nil {
        fmt.Println("List is empty")
        return
    }
    current := d.Tail
    for current != nil {
        fmt.Printf("%d ", current.Data)
        current = current.Prev
    }
    fmt.Println()
}

func main() {
    list := &DoublyLinkedList{}

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
