package main

import "fmt"

type Test interface {
	A()
}

type IShape interface {
	//Test
	Area()
	Perimeter()
}

type Traingle struct {
	side int
}

func (tra Traingle) A() {

	fmt.Println("A of Traingle is called.")
}
func (tra Traingle) Area() {

	fmt.Println("Area of Traingle is called.")
}

func (tra Traingle) Perimeter() {
	fmt.Println("Perimeter of Traingle is called.")
}

type Square struct {
	side int
}

func (sq Square) Area() {
	fmt.Println("Area of Square is called.")
}

func (sq Square) Perimeter() {
	fmt.Println("Perimeter of Square is called.")

}

func Calculate(shape IShape) {

	shape.Area()
	shape.Perimeter()
}

func log(t Test) {

}

func main() {

	var tra = Traingle{10}

	Calculate(tra)

	var sq = Square{20}

	Calculate(sq)

	log(tra)

	//var a Test = Traingle{30}

	//var shape IShape
	//When we try to assign object of struct to interface type of variable at that time it will check if that struct is implementing
	//that interface or not

	// shape = Traingle{side: 10}
	// shape.Area()
	// shape.Perimeter()

	// shape = Square{side: 20}

	// shape.Area()
	// shape.Perimeter()

}
