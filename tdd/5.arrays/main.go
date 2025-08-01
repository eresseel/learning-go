package main

import "fmt"

func GetArrays() ([3]int, [5]int) {
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	return arr1, arr2
}

func main() {
	arr1, arr2 := GetArrays()
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
}
