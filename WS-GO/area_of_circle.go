package main

import (
	"fmt"
)

// var (
// 	// a string = "Ram"
// 	b = 25
// 	c = 3.12
// 	my_Var =  "new"
// )
	

func main() {

	// b := "Raj"
	// aa,bb,cc := 22,"s",2.5
	const PI = 3.14
	var rad float64 = 2

	var Area float64 = (PI*rad*rad)

	fmt.Printf("Area of the circle is : %v and it's type is : %T \n",Area ,Area)
	fmt.Printf("Area of the circle is : %#v and it's type is : %T and %% ",Area ,Area)

}