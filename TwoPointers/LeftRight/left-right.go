package main

import "fmt"

// Two sum - find i, j -> arr[i] + arr[j] == target

func twoSum (arr []int, target int) []int {
	left, right := 0, len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{left+1, right+1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 3)) // arr ..., target = 3 -> 1+2 = target(3) Output -> [1 2]
}
