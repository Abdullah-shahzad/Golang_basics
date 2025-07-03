package main

import "fmt"

func main() {
    str := "Hi"
    bytes := []byte(str)  // convert string to byte slice
    fmt.Println(bytes)    // Output: [72 105]
    fmt.Println(string(bytes))  // Output: Hi
}
