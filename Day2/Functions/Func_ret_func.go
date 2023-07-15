package main

import "fmt"

func main() {

	//OuterFunc(10) : it is returning innerFUnc defn

	var exp = OuterFunc(5)

	fmt.Println(exp())

	fmt.Println(OuterFunc(10)())

	fmt.Println(OuterFunc(20)())

	fmt.Println(OuterFunc(30)())
}

func OuterFunc(id int) func() int {

	var counter = 0

	innerfunc := func() int {

		return id + counter
	}

	return innerfunc

}
