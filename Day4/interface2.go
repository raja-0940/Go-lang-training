package main




import "fmt"


type IShape interface {

    IDimen

    Area()

    Perimeter()

}


type IDimen interface {

    Topology()

}




type Traingle struct {

    side int

}




func (tra Traingle) Area() {

    fmt.Println("Area of Traingle is called.")

}




func (tra Traingle) Perimeter() {

    fmt.Println("Perimeter of Traingle is called.")

}




func (tra Traingle) Topology() {

    fmt.Println("Topology of Traingle is called.")

}




type Square struct {

    side int

}




func main() {

    var shape IShape

    shape = Traingle{side: 10}




    var dimen IDimen

    dimen = Traingle{side: 10}




    shape.Area()

    shape.Perimeter()

    dimen.Topology()

}