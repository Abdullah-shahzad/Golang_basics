package main
import "fmt"

func main() {
    var i interface{} = 7
    switch i.(type) {
    case int:
        fmt.Println("i is an int") // Output: i is an int
    case string:
        fmt.Println("i is a string")
    default:
        fmt.Println("unknown type")
    }
}