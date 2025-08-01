// main.go
package main

import "fmt"

func Slices() []int {
	myslice := []int{}
	return myslice
}

func Slices2() []string {
	myslice := []string{"Go", "Slices", "Are", "Powerful"}
	return myslice
}

func main() {
	fmt.Println("Slices() length:", len(Slices()))
	fmt.Println("Slices() capacity:", cap(Slices()))
	fmt.Println("Slices() values:", Slices())

	fmt.Println("Slices2() length:", len(Slices2()))
	fmt.Println("Slices2() capacity:", cap(Slices2()))
	fmt.Println("Slices2() values:", Slices2())
}
