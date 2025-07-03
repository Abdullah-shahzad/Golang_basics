package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }

//===============================================================

  x := 30
  if x >= 10 {
    fmt.Println("x is larger than or equal to 10.")
  } else if x > 20 {
    fmt.Println("x is larger than 20.")
  } else {
    fmt.Println("x is less than 10.")
  }

}