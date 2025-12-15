package main

import "fmt"

func binarySearch(arr [5]int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	arr := [5]int{15, 21, 33, 46, 57}
	target := 33

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Element - %d, Index - %d\n", target, index)
	} else {
		fmt.Printf("Element %d is not found\n", target)
	}
}
