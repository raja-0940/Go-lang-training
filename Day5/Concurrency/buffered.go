package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int, 3)

	ch1 <- 100
	ch1 <- 200
	ch1 <- 300

	fmt.Println("value of ch1 ", <-ch1)
	fmt.Println("value of ch1 ", <-ch1)

}
