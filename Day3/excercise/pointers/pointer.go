package main

import "fmt"

func main() {

	var i = 100

	// var ipointer *int
	// ipointer = &i
	// fmt.Println(ipointer)

	fmt.Println("Before calling changeValue : ", i)
	// changeValue(i)
	changeValue(&i)
	fmt.Println("After calling changeValue : ", i)  //100

}

func changeValue(value *int){
	
	*value = 999 // Accessing value from the memory address/pointer
	
}


// func changeValue(value int){
	
// 	value = 999
// 	fmt.Println(value)

// }