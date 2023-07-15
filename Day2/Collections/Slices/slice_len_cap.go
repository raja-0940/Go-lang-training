package main

import "fmt"

func main() {

	//len : current no of elements present in slice
	//cap : no of elements that can be hold by internal array

	slice1 := []int{10, 20, 30, 40, 50}

	fmt.Println(slice1)

	fmt.Println("Len of slice1 ", len(slice1))

	fmt.Println("Cap of slice1 ", cap(slice1))

	slice1 = append(slice1, 60)

	fmt.Println("Len of slice1 ", len(slice1))

	fmt.Println("Cap of slice1 ", cap(slice1))

	slice1[7] = 900

}
