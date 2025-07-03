package main

import (
	"fmt"
)


func main() {
  arr0 := [...]int{} 
  arr1 := [5]int{} 
  arr2 := [5]int{1,2} 
  arr3 := [5]int{1,2,3,4,5} 
  arr4 := [5]int{1:10,2:40}

  fmt.Println(arr0)
  fmt.Println(len(arr0))
  fmt.Println(arr1)
  fmt.Println(len(arr1))
  fmt.Println(arr2)
  fmt.Println(arr3)
  fmt.Println(arr4)
}


