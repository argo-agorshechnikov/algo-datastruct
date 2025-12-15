package main

import "fmt"

func main() {

	arr := [10]int{-1, 15, 2, 17, -999}

	for _, v := range arr {
		if v == 17 {
			fmt.Println(v)
			break
		}
	}

}
