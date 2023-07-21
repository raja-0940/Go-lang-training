package main

import (
	"fmt"
)

func main() {
	x := 4
	y := 7
	if x == 0 && y == 0 {
		fmt.Println("x and y are equals 0")
	} else if x < 5 && y < 5{
		fmt.Println("x and y are less than 5")
	} else {
		fmt.Println(" x and y are > 5")
	}

	switch x{
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
	case 3:
		fmt.Println("x=3")
	default:
		fmt.Println("x>3")
	}
}