package main

import "fmt"

// O(n^2)
func insertSort(arr [5]int) [5]int {
	
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	return arr
}

func main() {
	
	arr := [5]int {5, 4, 3, 2, 1}
	fmt.Println("Current array: ", arr)

	sortArr := insertSort(arr)
	fmt.Println("Sorted array: ", sortArr)

}
