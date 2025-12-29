package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i, j := start, end-1

	for i <= j {
		for i <= j && arr[i] < pivot {
			i++
		}

		for i <= j && arr[j] > pivot {
			j--
		}

		if i <= j {
			swap(arr, i, j)
			i++
			j--
		}
	}

	swap(arr, i, end)
	return i
}

func RecursiveQuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	swap(arr, mid, end)

	part := partition(arr, start, end)

	RecursiveQuickSort(arr, start, part)
	RecursiveQuickSort(arr, part+1, end)
}



func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 9, 6}
	fmt.Println(arr)

	fmt.Println("Recursive Quick Sort")
	RecursiveQuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
