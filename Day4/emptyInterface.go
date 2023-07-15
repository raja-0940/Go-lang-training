package main

import "fmt"

func main() {

	Print1("Persistent")
	Print1(10000)
	Print1([]int{10, 20, 30})

	array1  := []any {10,"Mukta",[]int{10,30,40}}

}

// func Print1(values ...interface{}){

// }

func Print1(value any) {
	fmt.Println(value)
}
