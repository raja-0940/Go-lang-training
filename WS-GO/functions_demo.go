package main

import (
	"fmt"
)

func helloWorld(){
	fmt.Println("Hello, world..!")
}

func familyName(fName string, age int) {
	fmt.Println("Name is Rajakumar",fName, "and his age is ",age)
}

// func add(a int, b int) int {
// 	var c = a + b
// 	return c
// }
func add(a int, b int) (c int) {
	c = a + b
	return
}

func userInfo(name string, age int, role string) (A int, name_and_role string) {
	A =  age
	name_and_role = name + role	
	return
}

// Recursion functions  -> call itself and reaches stop condition
func testcount(i int) int {
	if i == 11 {
		return 0
	}
	fmt.Println(i)
	return testcount(i+1)
}


func main() {
	helloWorld()
	familyName("Battula",32)
	fmt.Println(add(20,30))
	ans := add(30,60)
	fmt.Println(ans)	
	fmt.Println(userInfo("Rama",100,"Deva"))
	_,b := userInfo("Rama",100,"Deva") // used blank operator to omit its first return type
	fmt.Println(b)
	testcount(3)
	testcount(5)
	testcount(2)	
}