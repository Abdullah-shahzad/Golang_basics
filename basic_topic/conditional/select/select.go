package main

import (
	"fmt"
	"time"
)
func task_1(ch1 chan string){
	time.Sleep(2 * time.Second)
	ch1 <- "abdullah"

}
func task_2(ch2 chan string){
	time.Sleep(1 * time.Second)
	ch2 <- "shahzad"

}
func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	
	go task_1(ch1)
	
	go task_2(ch2)

	select{
	case m1 := <-ch1:
		fmt.Println(m1)
	case m2 := <-ch2:
		fmt.Println(m2)
	default:
		println("choudhary")
	}

}