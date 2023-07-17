package main

import "fmt"

func main() {
	var sl1 []int //nil slice
	fmt.Println(sl1)
	// sl1[0]=10 // we'll get error
	sl1=append(sl1,10,20,30)	
	fmt.Println(sl1)
}