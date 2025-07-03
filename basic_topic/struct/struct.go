package main
import("fmt")

type Person struct{
	id int 
	name string 
}


func main(){
	var p1 Person
	p1.id=12
	p1.name="Tim"

	var p2 Person
	p2.id=22
	p2.name="Tom"

	var p3 Person
	p3.id=32
	p3.name="Zeus"

	PrintPerson(p1 )
	PrintPerson(p2 )
	PrintPerson(p3 )

}

func PrintPerson(p Person){
	fmt.Println("ID: ",p.id) 
	fmt.Println("Name: ",p.name) 
	fmt.Println("------------------------") 
	
}