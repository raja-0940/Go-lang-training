package main

import (
	"fmt"
	"reflect"
)

type i12 interface {
}

//to empty interface we can pass any data type member

func findType(i interface{}) {

	
	//func findType(i i12) {
	
	fmt.Println(10,"abc",true)

	//fmt.Println(reflect.TypeOf(i))
	//fmt.Println(i.(type))

	switch i.(type) { // type assertion

	case string:
		fmt.Println(i.(string))
		var value = i.(string)
		fmt.Printf("%T\n", value)
	case int:
		fmt.Println(i.(int))

	case Product:
		var value1 = i.(Product)

		fmt.Printf("%T\n", value1)

	}
}

func myFunc(param ...interface{}) {

}

type Product struct {
}

func main() {
	myFunc(10, "My Name", 190.45)
	findType(100)
	findType("Mukta")
	findType(Product{})

	fmt.Println(10,"abc",true)
}
