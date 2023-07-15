package main

import "fmt"

// global var
var gvar int

func main() {
	//local vars
	var name string = "Raja"
	var age int = 29
	var bstatus bool = true

	fmt.Println(name, age, bstatus)

	var FName, LName = "Raja", "Kumar"
	fmt.Println(FName,LName)
}