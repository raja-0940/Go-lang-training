package main

import (
	"fmt"
	"reflect"
)

func main() {

	//slice1 := make([]int,2) //length & capacity both are 2

	slice1 := make([]int, 2, 4)

	fmt.Println(reflect.TypeOf(slice1))

	slice1[0] = 100
	slice1[1] = 200
	fmt.Println(slice1)

	fmt.Println("Len of slice1 ", len(slice1))

	fmt.Println("Cap of slice1 ", cap(slice1))

	slice2 := append(slice1, 300, 400, 500)

	slice1[0] = 777777

	fmt.Println(slice1)

	fmt.Println(slice2)

}
