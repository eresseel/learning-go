package main

import ("fmt")

func main() {
  prices := []int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])

  // Change Elements of a Slice
  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])

  // Append Elements To a Slice
  // slice_name = append(slice_name, element1, element2, ...)
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // Append One Slice To Another Slice
  myslice2 := []int{1,2,3}
  myslice3 := []int{4,5,6}
  myslice4 := append(myslice2, myslice3...)

  fmt.Printf("myslice4=%v\n", myslice4)
  fmt.Printf("length=%d\n", len(myslice4))
  fmt.Printf("capacity=%d\n", cap(myslice4))

  // Change The Length of a Slice
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice5 := arr1[1:5] // Slice array
  fmt.Printf("myslice5 = %v\n", myslice5)
  fmt.Printf("length = %d\n", len(myslice5))
  fmt.Printf("capacity = %d\n", cap(myslice5))

  myslice5 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice5 = %v\n", myslice5)
  fmt.Printf("length = %d\n", len(myslice5))
  fmt.Printf("capacity = %d\n", cap(myslice5))

  myslice5 = append(myslice5, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice5 = %v\n", myslice5)
  fmt.Printf("length = %d\n", len(myslice5))
  fmt.Printf("capacity = %d\n", cap(myslice5))

  // Memory Efficiency
  // copy(dest, src)
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
}