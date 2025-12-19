package main

import "fmt"

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))

    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    // Добавляем оставшиеся элементы
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}

func main() {
    arr := []int{4, 2, 5, 1, 3}
    fmt.Println("Исходный массив:", arr)
    sorted := mergeSort(arr)
    fmt.Println("Отсортированный массив:", sorted)
}
