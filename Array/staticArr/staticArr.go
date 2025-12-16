package main

import "fmt"

func arrIteration(arr [5]int) {

	for i, v := range arr {
		fmt.Printf("Index - %d, val - %d\n", i, v)
	}
}

func main() {
	// O(n)
	arr := [5]int{1, 2, 3, 4, 5}

	// O(1)
	fmt.Println(arr[1])
	fmt.Println("---------------------------------------------")

	//O(1)
	arr[1] = 0
	fmt.Println(arr[1])
	fmt.Println("---------------------------------------------")

	// O(n)
	arrIteration(arr)
	fmt.Println("---------------------------------------------")

	// O(n)
	newArr := arr
	fmt.Println(newArr)

}
