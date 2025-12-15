package main

import "fmt"

// Динамический двумерный массив

type DynamicMatrix struct {
	data [][]int
}

func NewDynamicMatrix() *DynamicMatrix {
	return &DynamicMatrix{
		data: make([][]int, 0),
	}
}

func (dm *DynamicMatrix) AppendRow(row []int) {
	newData := make([][]int, len(dm.data)+1)

	for i := 0; i < len(dm.data); i++ {
		newData[i] = dm.data[i]
	}

	newData[len(dm.data)] = row
	dm.data = newData
}

func (dm *DynamicMatrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(dm.data) || col < 0 || col >= len(dm.data[row]) {
		return false
	}
	dm.data[row][col] = val
	return true
}

func (dm *DynamicMatrix) Get(row, col int) (int, bool) {
	if row < 0 || row >= len(dm.data) || col < 0 || col >= len(dm.data[row]) {
		return 0, false
	}
	return dm.data[row][col], true
}

func (dm *DynamicMatrix) Print() {
	for _, row := range dm.data {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func main() {
	dm := NewDynamicMatrix()
	dm.AppendRow([]int{1, 2, 3})
	dm.AppendRow([]int{4, 5, 6})
	dm.AppendRow([]int{7, 8, 9, 10})

	dm.Print()

	fmt.Println()

	dm.Set(0, 1, 20)
	dm.Print()
}
