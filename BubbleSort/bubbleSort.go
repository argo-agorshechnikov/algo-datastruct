package main

import "fmt"

func bubbleSort(arr [5]int) [5]int {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	arr := [5]int{122, 1, 234, 51, 19}
	fmt.Println("Current arr: ", arr)

	sortArr := bubbleSort(arr)
	fmt.Println("Sorted arr: ", sortArr)
}
