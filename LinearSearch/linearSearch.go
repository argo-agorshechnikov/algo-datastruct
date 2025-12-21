package main

import "fmt"

func main() {

	arr := [10]int{-1, 15, 2, 17, -999}

	// O(n)
	target := 2
	for i := range arr {
		if i == target {
			fmt.Println(arr[target])
			break
		}
	}
	
}
