package main

import (
	"fmt"
)

func myFunction1(x int, y int) int {
	return x + y
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction3(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction4(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction1(1, 2))
	fmt.Println(myFunction2(1, 2))
	total := myFunction3(1, 2)
	fmt.Println(total)
	fmt.Println(myFunction4(5, "Hello"))
	_, b := myFunction4(5, "Hello")
	fmt.Println(b)
	a, _ := myFunction4(5, "Hello")
	fmt.Println(a)
}
