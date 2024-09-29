package main

import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)

  // Create a Slice From an Array
  // var myarray = [length]datatype{values} // An array
  // myslice := myarray[start:end] // A slice made from the array
  arr1 := [6]int{10, 11, 12, 13, 14, 15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))

  // Create a Slice With The make() Function
  // slice_name := make([]type, length, capacity)
  myslice3 := make([]int, 5, 10)
  fmt.Printf("myslice3 = %v\n", myslice3)
  fmt.Printf("length = %d\n", len(myslice3))
  fmt.Printf("capacity = %d\n", cap(myslice3))

  // with omitted capacity
  myslice4 := make([]int, 5)
  fmt.Printf("myslice4 = %v\n", myslice4)
  fmt.Printf("length = %d\n", len(myslice4))
  fmt.Printf("capacity = %d\n", cap(myslice4))
}