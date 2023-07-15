package main

import "fmt"

func main(){
	// exp := outerfunction()
	//exp will have defination of inner function which returns by outer function
	// fmt.Println(exp())
	fmt.Println(outerfunction(10)())
	fmt.Println(outerfunction(20)())
	fmt.Println(outerfunction(30)())
}

func outerfunction(ID int) func() int{
	innerfunction := func() int {
		return ID + 10
	}
	return innerfunction
}