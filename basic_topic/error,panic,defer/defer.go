package main

import("fmt")

/*
Defer
*/

// func main(){
// 	defer fmt.Println("1")
// 	fmt.Println("2")
// 	defer fmt.Println("3")
// }


/*
Panic
*/

// func main(){
// 	defer func(){
// 	if e:=recover();e!=nil {
// 		fmt.Println("Recovered from panic")
// 	}
// }()

// 	panic("Random panic error")
// }


/*
Error handling
*/

func main(){
	var num int 

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("We got an error!")
		fmt.Println(err)
	}else{
	fmt.Println(num)
	}
}