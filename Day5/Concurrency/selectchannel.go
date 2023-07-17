package main

import (
	"fmt"
	"time"
)

func Server1(ch1 chan string) {

	time.Sleep(4 * time.Second)
	ch1 <- " Response from Server1"
}

func Server2(ch2 chan string) {

	time.Sleep(2 * time.Second)
	ch2 <- " Response from Server2"

}

func main() {

	ch1 := make(chan string)

	ch2 := make(chan string)
	go Server1(ch1)
	go Server2(ch2)
	// we r going to call both go routines & will wait till we get response from either of these go routine
	//when we r going to listen for a data from multiple channels we will use select
	select {
	case value := <-ch1:
		fmt.Println(value)
	case value := <-ch2:
		fmt.Println(value)
	}
}
