package main

import "fmt"

func matrixIteration (dMatrix [][]int) {

	for i := 0; i < len(dMatrix); i++{
		for j := 0; j < len(dMatrix[i]); j++{
			fmt.Print(dMatrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func addRow(matrix [][]int, row []int) [][]int {
	return append(matrix, row)
}

func addColumn(matrix [][]int, col []int) [][]int{
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append(matrix[i], col[i])
	}

	return matrix
}

func removeEndRow(matrix [][]int, index int) [][]int {
	return append(matrix[:index], matrix[index+1:]...)
}

func removeCol(matrix [][]int, index int) [][]int {
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append(matrix[i][:index], matrix[i][index+1:]...)
	}

	return matrix
}

func insertRow(matrix [][]int, row []int, index int) [][]int {
	if index > len(matrix) {
		index = len(matrix)
	}

	matrix = append(matrix, nil)
    copy(matrix[index+1:], matrix[index:])
    matrix[index] = row

    return matrix
}

func insertColumn(matrix [][]int, col []int, index int) [][]int {
    for i := 0; i < len(matrix); i++ {
        if index > len(matrix[i]) {
            matrix[i] = append(matrix[i], col[i])
        } else {
            matrix[i] = append(matrix[i][:index], append([]int{col[i]}, matrix[i][index:]...)...)
        }
    }
    return matrix
}

func removeAnyRow(matrix [][]int, index int) [][]int {
    return append(matrix[:index], matrix[index+1:]...)
}

func removeAnyColumn(matrix [][]int, index int) [][]int {
    for i := 0; i < len(matrix); i++ {
        if index < len(matrix[i]) {
            matrix[i] = append(matrix[i][:index], matrix[i][index+1:]...)
        }
    }
    return matrix
}



func main() {
	dynamicMatrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8, 9, 10},
	}

	// O(1) Получение элемента по индексу
	fmt.Println(dynamicMatrix[0][4])
	fmt.Println("----------------------------------------------")

	// O(1) Изменение элемента по индексу
	dynamicMatrix[2][2] = 100
	fmt.Println(dynamicMatrix[2][2])
	fmt.Println("----------------------------------------------")

	// O(m*n) m - rows, n - cols (Итерация по матрице)
	matrixIteration(dynamicMatrix)
	fmt.Println("----------------------------------------------")

	// O(m*n) m - rows, n - cols (Копирование матрицы)
	newDMatrix := dynamicMatrix
	fmt.Printf("newDMatrix - %v\n", newDMatrix)
	fmt.Println("----------------------------------------------")

	// O(1) Добавление строки в конец
	row := []int{111, 222, 333}
	fmt.Println(dynamicMatrix)
	fmt.Println(addRow(dynamicMatrix, row))
	fmt.Println("----------------------------------------------")

	// O(n) Добавление столбца в конец
	col := []int{0, 0, 0}
	fmt.Println(dynamicMatrix)
	fmt.Println(addColumn(dynamicMatrix, col))
	fmt.Println("----------------------------------------------")

	// O(1) Удаление строки с конца
	fmt.Println(dynamicMatrix)
	fmt.Println(removeEndRow(dynamicMatrix, 2))
	fmt.Println("----------------------------------------------")

	// O(m*n) m - rows, n - cols (Удаление столбцов с индексом 2 из всей матрицы)
	fmt.Println(dynamicMatrix)
	fmt.Println(removeCol(dynamicMatrix, 2))
	fmt.Println("----------------------------------------------")

	// O(n) n - rows (Добавление строки в произвольное место)
	fmt.Println(dynamicMatrix)
	insrtRow := []int{1, 1, 1}
	fmt.Println(insertRow(dynamicMatrix, insrtRow, 1))
	fmt.Println("----------------------------------------------")

	// O(m*n) m - rows, n - cols (Добавление столбца в произвольное место)
	fmt.Println(dynamicMatrix)
	insrtCol := []int{999, 999, 999}
	fmt.Println(insertColumn(dynamicMatrix, insrtCol, 2))
	fmt.Println("----------------------------------------------")

	// O(n) n - rows (Удаление строки из произвольного места)
	fmt.Println(dynamicMatrix)
	fmt.Println(removeAnyRow(dynamicMatrix, 1))
	fmt.Println("----------------------------------------------")

	// O(m*n) m - rows, n - cols (Удаление столбца из произвольного места)
	fmt.Println(dynamicMatrix)
	fmt.Println(removeAnyColumn(dynamicMatrix, 2))

}