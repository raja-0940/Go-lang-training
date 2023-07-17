package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice1 := make([]int,2,4) //length is 2 and capacity is 4

	fmt.Println(reflect.TypeOf(slice1))

	slice1[0]=100
	slice1[1]=200
	fmt.Println(slice1)
	fmt.Println("Lengthof slice",len(slice1))
	fmt.Println("Capacity of slice",cap(slice1))

	slice2:=append(slice1,300,400,500,600)
	slice1[0]=9999
	fmt.Println(slice1)
	fmt.Println(slice2)
}