package main

import (
	"fmt"
)

// package imported & not used

func main() {

	var a, b = 100, 200

	fmt.Println(b)

	_ = a // for variables declared & not used

}
