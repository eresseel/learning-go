// Print()
// Println()
// Printf()

package main

import ("fmt")

func main() {
  var i,j string = "Hello","World"
  var k int = 15

  fmt.Print(i)
  fmt.Print(j)
  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
  fmt.Print(i, "\n",j)
  fmt.Print(i, " ", j)

  fmt.Println(i,j)

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("k has value: %v and type: %T", k, k)
}