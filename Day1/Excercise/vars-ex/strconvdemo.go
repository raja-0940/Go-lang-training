package main

import (
	"fmt"
	"strconv"
)

func main(){
	a := "10"
	// fmt.Println(a)
	// value, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(value)
	// }
	value, _ := strconv.Atoi(a)  //string to int conversion
	fmt.Println(value)

	b := 20
	value1 := strconv.Itoa(b)  //int to string conversion
	fmt.Println(value1)
	
}