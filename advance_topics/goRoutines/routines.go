package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello() // This runs in the background

    time.Sleep(1 * time.Second) // Wait so goroutine can finish
    fmt.Println("Main function done.")
}