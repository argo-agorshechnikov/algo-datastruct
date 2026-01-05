package main

import "fmt"

func heapSort(arr []int) {

	buildeMaxHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		heapifyDown(arr, 0, i)
	}

}

func buildeMaxHeap(arr []int) {

	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapifyDown(arr, i, len(arr))
	}
}

func heapifyDown(arr []int, i, n int) {
	for {
		largest := i     // текущий максимум
		left := 2*i + 1  // левый ребенок
		right := 2*i + 2 // правый ребенок

		// проверяем левого ребенка
		if left < n && arr[left] > arr[largest] {
			largest = left
		}

		// проверяем правого ребенка
		if right < n && arr[right] > arr[largest] {
			largest = right
		}

		// если текущий — максимум, выходим
		if largest == i {
			break
		}

		// swap и продолжаем вниз
		swap(arr, i, largest)
		i = largest
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{3, 1, 4, 5, 2}
	fmt.Println("Current arr: ", arr)

	heapSort(arr)
	fmt.Println("Sorted arr: ", arr)
}
