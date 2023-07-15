package main

import "fmt"

var funcexp = func() int {
	//fmt.Println("Inside funcExpression")
	return 1000
}

func main() {

	fmt.Println(funcexp())

	var exp = func() {

		fmt.Println("Hello From greet Function")
	}

	exp()

	// Hello()

}

func Hello() {} // 1 way of defining function (FUnction defn)
