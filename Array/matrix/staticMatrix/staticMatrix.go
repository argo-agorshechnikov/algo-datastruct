package main

import "fmt"

func matrixIteration (matrix [3][3]int) {
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++{
			fmt.Print(matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	staticMatrix := [3][3]int{
		{1, 2, 3}, 
		{4, 5, 6}, 
		{7, 8, 9},
	}

	// O(1)
	fmt.Println(staticMatrix[1][1])
	fmt.Println("-------------------------------------------------------")

	// O(1)
	staticMatrix[1][1] = 555
	fmt.Println(staticMatrix)
	fmt.Println("-------------------------------------------------------")

	// O(m*n) m - rows, n - cols
	matrixIteration(staticMatrix)
	fmt.Println("-------------------------------------------------------")

	// O(m*n) m - rows, n - cols
	newMatrix := staticMatrix
	fmt.Printf("newM %v\n", newMatrix)
}