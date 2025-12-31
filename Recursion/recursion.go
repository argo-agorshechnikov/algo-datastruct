package main

import "fmt"

func fact(n int) int {
	if n <= 1 {
		return 1 // basic case
	}

	return n * fact(n - 1) // recursive case
}

func main() {
	fmt.Println(fact(5))
}
