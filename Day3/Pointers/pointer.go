package main

import "fmt"

func main() {

	var i = 100

	// var iPointer *int

	// iPointer = &i

	// fmt.Println(iPointer)

	fmt.Println("Before calling changeValue ", i) // 100

	changeValue(&i)

	fmt.Println("After calling changeValue ", i) //

}

func changeValue(value *int) {

	*value = 999 // accessing value from memory address/pointer
}

// func changeValue(value int) {

// 	value = 999
// }
