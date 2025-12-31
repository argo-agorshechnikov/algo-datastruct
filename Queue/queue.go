package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	
	return item, true
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func main() {
	q := &Queue[int]{}
    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)

	fmt.Println(*q)
    
	if item, ok := q.Dequeue(); ok {
		fmt.Println("Dequeue: ", item)
	}

	fmt.Println(*q)

	fmt.Println("Size: ", q.Size())
}