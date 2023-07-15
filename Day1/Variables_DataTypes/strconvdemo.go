package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i = "1000abc"

	value, err := strconv.Atoi(i) // string to int conversion

	fmt.Println(value, err)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {

	// 	fmt.Println(value)
	// }

	// var j = 2000

	// var value1 = strconv.Itoa(j)

	// fmt.Println(value1)

}
