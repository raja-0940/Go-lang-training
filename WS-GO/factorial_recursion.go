package main

import (
	"fmt"
)

func factorial(a float64) (b float64) {
	if a > 0 {
		b = a * factorial(a-1)
	}else{
		b = 1
	}
	return
}

func main() {
	fact := factorial(5)
	fmt.Println
}