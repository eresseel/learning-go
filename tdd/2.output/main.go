package main

import (
	"fmt"
)

func ValuePI() float32 {
	const PI = 3.14

	return PI
}

func main() {
	var i string = "The value of PI: "
	fmt.Println(i, ValuePI())
}
