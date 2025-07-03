package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

//recurrsive function

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}


func main() {


  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)

  testcount(1)

}