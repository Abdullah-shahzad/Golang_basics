package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }

  for j:=0; j < 5; j++ {
    if j == 3 {
      continue
    }
   fmt.Println(j)
  }

  for k:=0; k < 5; k++ {
    if k== 3 {
      break
    }
   fmt.Println(k)
  }



// unconditional loop with range keyword


  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
  
  for _, val := range fruits {
     fmt.Printf("%v\n", val)
  }


  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }

}