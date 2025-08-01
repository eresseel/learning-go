package main

import "fmt"

func TestCount(x int) []int {
	var result []int

	for i := x; i < 11; i++ {
		result = append(result, i)
	}

	return result
}

func main() {
	for _, v := range TestCount(1) {
		fmt.Println(v)
	}
}
