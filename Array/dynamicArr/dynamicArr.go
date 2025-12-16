package main

import "fmt"

func darrInteration(darr []int) {

	for i, v := range darr {
		fmt.Printf("Index - %d, value - %d\n", i, v)
	}
}

func deleteValue(darr []int, index int) []int {

	return append(darr[:index], darr[index+1:]...)
}

func insertValue(darr []int, index, value int) []int {
	darr = append(darr, 0) // [1, 2, 3, 4, 5, 0]

	copy(darr[index + 1:], darr[index:]) // index == 2(val-3) [3, 4, 5, 0] copy to [4, 5, 0] => [1, 2, 3, 3, 4, 5]

	darr[index] = value

	return darr
}

func main() {
	// O(n)
	dynamicArr := []int{1, 2, 3, 4, 5}

	// O(1)
	fmt.Println(dynamicArr[1])
	fmt.Println("---------------------------------------------")

	// O(1)
	dynamicArr[4] = 555
	fmt.Println(dynamicArr[4])
	fmt.Println("---------------------------------------------")

	// O(n)
	darrInteration(dynamicArr)
	fmt.Println("---------------------------------------------")

	// O(n)
	newDArr := dynamicArr
	fmt.Println(newDArr)
	fmt.Println("---------------------------------------------")

	// O(1) - в среднем, O(n) - при расширении
	dynamicArr = append(dynamicArr, 666)
	fmt.Println(dynamicArr)
	fmt.Println("---------------------------------------------")

	// O(1)
	dynamicArr = deleteValue(dynamicArr, 5)
	fmt.Println(dynamicArr)
	fmt.Println("---------------------------------------------")

	// O(n)
	dynamicArr = deleteValue(dynamicArr, 2)
	fmt.Println(dynamicArr)
	fmt.Println("---------------------------------------------")

	// O(n)
	fmt.Println(insertValue(dynamicArr, 2, 333))
	fmt.Println("---------------------------------------------")


}
