package main

import "fmt"

var funcexp = func() int{
	return 1000
}

func main(){

	fmt.Println(funcexp()) //calling anonymous function with return int value

	var exp = func(){
		fmt.Println("Hello from anonymous function")
	}
	fmt.Println(exp) //return memory location of exp 
	exp() // calling anonymous function

	hello() // simple function calling
	
}

func hello(){
	fmt.Println("Hello from simple function")
}