package main

import "fmt"

type DynamicArr struct {
	data []int
	length int
	capacity int
}

func NewDArr(startCapacity int) *DynamicArr {
	return &DynamicArr{
		data: make([]int, startCapacity),
		length: 0,
		capacity: startCapacity,
	}
}

func (da *DynamicArr) resize() {
	newCapacity := da.capacity * 2
	newData := make([]int, newCapacity)
	copy(newData, da.data)
	da.data = newData
	da.capacity = newCapacity
}

func (da *DynamicArr) Append(val int) {
	if da.length >= da.capacity {
		da.resize()
	}

	da.data[da.length] = val
	da.length++
}

func (da *DynamicArr) Get(i int) (int, bool) {
	if i < 0 || i >= da.length {
		return 0, false
	} 

	return da.data[i], true
}

func (da *DynamicArr) Set(i int, val int) bool {
	if i < 0 || i >= da.length {
		return false
	}

	da.data[i] = val
	return true 
}

func (da *DynamicArr) Length() int{
	return da.length
}

func (da *DynamicArr) Capacity() int {
	return da.capacity
}

func main() {
	arr := NewDArr(2)
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)

	for i := 0; i < arr.Length(); i++ {
		v, _ := arr.Get(i)
		fmt.Println(v)
	}

}