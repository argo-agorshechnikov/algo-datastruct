package main

import "fmt"

// Статический одномерный массив

type StaticArr struct {
	data   [5]int
	length int
}

func NewStaticArr() *StaticArr {
	return &StaticArr{
		data:   [5]int{},
		length: 0,
	}
}

func (sa *StaticArr) Append(val int) bool {
	if sa.length >= len(sa.data) {
		return false
	}

	sa.data[sa.length] = val
	sa.length++
	return true
}

func (sa *StaticArr) Get(i int) (int, bool) {
	if i < 0 || i >= sa.length {
		return 0, false
	}

	return sa.data[i], true
}

func (sa *StaticArr) Set(i int, val int) bool {
	if i < 0 || i >= sa.length {
		return false
	}

	sa.data[i] = val
	return true
}

func (sa *StaticArr) Length() int {
	return sa.length
}

func main() {
	arr := NewStaticArr()
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)
	arr.Append(4)
	arr.Append(5)

	for i := 0; i < arr.Length(); i++ {
		v, _ := arr.Get(i)
		fmt.Printf("index - %d, val - %d\n", i, v)
	}

	fmt.Println()

	arr.Append(6)
	for i := 0; i < arr.Length(); i++ {
		v, _ := arr.Get(i)
		fmt.Printf("index - %d, val - %d\n", i, v)
	}
}
