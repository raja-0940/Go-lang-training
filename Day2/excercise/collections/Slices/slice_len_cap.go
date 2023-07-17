package main

import "fmt"

//len : current no of elements presented in a slice.
//cap : no of elements that can be hold by internal array
func main() {
	slice := []int{10,20}
	fmt.Println("Lengthof slice",len(slice))
	fmt.Println("Capacity of slice",cap(slice))

	slice=append(slice,30,40,50)
	fmt.Println("Lengthof slice",len(slice))
	fmt.Println("Capacity of slice",cap(slice))

	slice=append(slice,60,70,80,90)
	fmt.Println("Lengthof slice",len(slice))
	fmt.Println("Capacity of slice",cap(slice))

	slice[0]=100 // replaces current value with given value
	fmt.Println(slice)

}