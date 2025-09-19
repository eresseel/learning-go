package integers

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("The sum of number: ", Add(2, 2))
}
