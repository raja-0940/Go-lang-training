package main

import "fmt"

func main() {

	var slice1 = []int{10, 20, 30, 40}

	fmt.Println(slice1[0:3]) //slice[L:H].It will take elements from L to H-1

	fmt.Println(slice1[:2]) // it will take elements from 0 to H-1

	fmt.Println(slice1[1:]) //it will be from 1 to last element

}
