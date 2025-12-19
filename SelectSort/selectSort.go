package main

import "fmt"

// O(n^2)
func selectSort(arr [5]int) [5]int {

	for i := 0; i < len(arr)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }

	return arr
}

func main() {
	
	arr := [5]int {5, 3, 2, 1, 4}
	fmt.Println("Current arr: ", arr)

	sortArr := selectSort(arr)
	fmt.Println("Sorted arr: ", sortArr)
}
