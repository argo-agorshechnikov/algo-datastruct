package main

import "fmt"

type Stack[T any] struct {
	items []T
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, 16),
		size: 0,
	}
}

func (s *Stack[T]) Push(item T) {
	if s.size == len(s.items) {
		s.items = append(s.items, item)
	} else {
		s.items[s.size] = item
	}
	
	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	s.size--
	return s.items[s.size], true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	return s.items[s.size-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Clear() {
	s.size = 0
}


func main() {
	stack := NewStack[int]()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	fmt.Println("Size", stack.Size())

	if top, ok := stack.Peek(); ok {
		fmt.Println("Top", top)
	}

	if popped, ok := stack.Pop(); ok {
		fmt.Println("Pop", popped)
	}

	stack.Clear()
	fmt.Println("IsEmpty", stack.IsEmpty())
}
