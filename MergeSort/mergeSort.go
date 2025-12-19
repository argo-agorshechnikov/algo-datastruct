package main

import "fmt"

type Node struct {
    Data int
    Next *Node
}

func mergeSort(head *Node) *Node {
    if head == nil || head.Next == nil {
        return head
    }

    // Находим середину списка
    mid := findMiddle(head)
    left := head
    right := mid.Next
    mid.Next = nil

    // Рекурсивно сортируем обе половины
    left = mergeSort(left)
    right = mergeSort(right)

    // Сливаем отсортированные половины
    return merge(left, right)
}

func findMiddle(head *Node) *Node {
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func merge(left, right *Node) *Node {
    dummy := &Node{}
    current := dummy

    for left != nil && right != nil {
        if left.Data <= right.Data {
            current.Next = left
            left = left.Next
        } else {
            current.Next = right
            right = right.Next
        }
        current = current.Next
    }

    if left != nil {
        current.Next = left
    } else {
        current.Next = right
    }

    return dummy.Next
}

func printList(head *Node) {
    for head != nil {
        fmt.Printf("%d ", head.Data)
        head = head.Next
    }
    fmt.Println()
}

func main() {
    // Создаём список: 4 -> 2 -> 5 -> 1 -> 3
    head := &Node{Data: 4}
    head.Next = &Node{Data: 2}
    head.Next.Next = &Node{Data: 5}
    head.Next.Next.Next = &Node{Data: 1}
    head.Next.Next.Next.Next = &Node{Data: 3}

    fmt.Println("Исходный список:")
    printList(head)

    head = mergeSort(head)

    fmt.Println("Отсортированный список:")
    printList(head)
}
