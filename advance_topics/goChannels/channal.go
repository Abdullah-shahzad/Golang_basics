package main

import (
    "fmt"
)

func sendValue(ch chan string) {
    ch <- "Hello from channel!"
}

func main() {
    ch := make(chan string)

    go sendValue(ch)

    msg := <-ch // waits to receive value
    fmt.Println(msg)
}
