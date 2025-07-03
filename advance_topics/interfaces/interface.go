package main
import("fmt")

type Shape interface{
	area() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length,width float64
}

func (c Circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64{
	return r.length * r.width
}


func main(){
	var s Shape

	s = Circle{radius: 2}
	a := s.area()
	fmt.Println("Circle area = ",a)

	s = Rectangle{length: 2,width: 2}
	r := s.area()
	fmt.Println("Rectangle area = ",r)



	

}
