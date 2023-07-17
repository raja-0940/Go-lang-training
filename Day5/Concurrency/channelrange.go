package main

import "fmt"

func Producer(ch1 chan int) {

	for i := 0; i < 10; i++ {

		ch1 <- i
	}

	close(ch1)
}

func main() {

	ch1 := make(chan int)

	go Producer(ch1)

	for value := range ch1 {

		fmt.Println("Value from channel", value)
	}

}
