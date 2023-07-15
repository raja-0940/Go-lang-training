package main

import (
	"fmt"
)

func main() {

	// var value = test()

	// if value == 100 {

	// 	fmt.Println(value)
	// }

	if value := test(); value == 100 {
		fmt.Println("if is called")
	}
}

func test() int {

	return 100
}
