package main

import "fmt"

func main() {

	//Self Executing (IIFE) functions
	//It is called at aplace where it is defined
	//it can be called only once
	//Generally it is used for one time initilaization of memebers
	//ABstracting some data from outside world

	// func() {
	// 	fmt.Println("Self Executing function is called")
	// }() // We are calling this func here itself

	var exp = func() int {

		return 100
	}() //self executing function

	fmt.Println(exp)

	fmt.Println(exp)

	fmt.Println(exp)

}
