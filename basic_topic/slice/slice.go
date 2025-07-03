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

  // thorught make()

  myslice3 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice3)
  fmt.Printf("length = %d\n", len(myslice3))
  fmt.Printf("capacity = %d\n", cap(myslice3))

  // with omitted capacity
  myslice4 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice4)
  fmt.Printf("length = %d\n", len(myslice4))
  fmt.Printf("capacity = %d\n", cap(myslice4))
}
 

// modified or appended slice

// func main() {
//   myslice1 := []int{1, 2, 3, 4, 5, 6}
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21)
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))
// }
