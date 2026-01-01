package main

import (
	"fmt"
)

type MaxHeap []int
type MinHeap []int

// ===== Общие вспомогательные функции индексов =====

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

// ===================== MAX HEAP =====================

// siftUp для MaxHeap (всплытие)
func (h *MaxHeap) siftUp(i int) {
	for i > 0 && (*h)[i] > (*h)[parent(i)] {
		(*h)[i], (*h)[parent(i)] = (*h)[parent(i)], (*h)[i]
		i = parent(i)
	}
}

// siftDown для MaxHeap (погружение)
func (h *MaxHeap) siftDown(i int) {
	n := len(*h)
	for {
		l, r := left(i), right(i)
		largest := i

		if l < n && (*h)[l] > (*h)[largest] {
			largest = l
		}
		if r < n && (*h)[r] > (*h)[largest] {
			largest = r
		}
		if largest == i {
			break
		}
		(*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
		i = largest
	}
}

// Push добавляет элемент в MaxHeap
func (h *MaxHeap) Push(x int) {
	*h = append(*h, x)
	h.siftUp(len(*h) - 1)
}

// Pop извлекает максимум из MaxHeap
func (h *MaxHeap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	maxVal := (*h)[n]
	*h = (*h)[:n]
	if len(*h) > 0 {
		h.siftDown(0)
	}
	return maxVal, true
}

// Heapify строит MaxHeap из массива за O(n)
func (h *MaxHeap) Heapify(arr []int) {
	*h = append((*h)[:0], arr...) // копия, чтобы не портить исходный слайс
	for i := len(*h)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

// ===================== MIN HEAP =====================

// siftUp для MinHeap (всплытие)
func (h *MinHeap) siftUp(i int) {
	for i > 0 && (*h)[i] < (*h)[parent(i)] {
		(*h)[i], (*h)[parent(i)] = (*h)[parent(i)], (*h)[i]
		i = parent(i)
	}
}

// siftDown для MinHeap (погружение)
func (h *MinHeap) siftDown(i int) {
	n := len(*h)
	for {
		l, r := left(i), right(i)
		smallest := i

		if l < n && (*h)[l] < (*h)[smallest] {
			smallest = l
		}
		if r < n && (*h)[r] < (*h)[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		i = smallest
	}
}

// Push добавляет элемент в MinHeap
func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
	h.siftUp(len(*h) - 1)
}

// Pop извлекает минимум из MinHeap
func (h *MinHeap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	minVal := (*h)[n]
	*h = (*h)[:n]
	if len(*h) > 0 {
		h.siftDown(0)
	}
	return minVal, true
}

// Heapify строит MinHeap из массива за O(n)
func (h *MinHeap) Heapify(arr []int) {
	*h = append((*h)[:0], arr...)
	for i := len(*h)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

// ===================== DEMO =====================

func main() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	// MaxHeap
	var maxH MaxHeap
	maxH.Heapify(arr)
	fmt.Println("MaxHeap internal:", maxH)
	for len(maxH) > 0 {
		v, _ := maxH.Pop()
		fmt.Printf("Max pop: %d\n", v)
	}

	fmt.Println()

	// MinHeap
	var minH MinHeap
	minH.Heapify(arr)
	fmt.Println("MinHeap internal:", minH)
	for len(minH) > 0 {
		v, _ := minH.Pop()
		fmt.Printf("Min pop: %d\n", v)
	}
}
