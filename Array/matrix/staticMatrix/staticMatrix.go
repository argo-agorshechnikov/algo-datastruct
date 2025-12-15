package main

import "fmt"

// Статический двумерный массив

type StaticMatrix struct {
	data [3][3]int
	rows int
	cols int
}

func NewStaticMatrix() *StaticMatrix {
	return &StaticMatrix{
		data: [3][3]int{},
		rows: 3,
		cols: 3,
	}
}

func (sm *StaticMatrix) Set(row, col, val int) bool {

	if row < 0 || row >= sm.rows || col < 0 || col >= sm.cols {
		return false
	}

	sm.data[row][col] = val
	return true
}

func (sm *StaticMatrix) Get(row, col int) (int, bool) {
	if row < 0 || row >= sm.rows || col < 0 || col >= sm.cols {
		return 0, false
	}

	return sm.data[row][col], true
}

func (sm *StaticMatrix) Print() {
	for i := 0; i < sm.rows; i++ {
		for j := 0; j < sm.cols; j++ {
			v, _ := sm.Get(i, j)
			fmt.Print(v)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	sMatrix := NewStaticMatrix()
	sMatrix.Set(0, 0, 5)
	sMatrix.Set(1, 1, 55)
	sMatrix.Set(2, 2, 555)

	sMatrix.Print()
}
