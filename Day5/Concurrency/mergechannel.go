package main

import (
	"fmt"
	"time"
)

func Database1(ch1 chan string) {

	time.Sleep(4 * time.Second)
	ch1 <- " Response from Database1"
}

func Database2(ch2 chan string) {

	time.Sleep(2 * time.Second)
	ch2 <- " Response from Database2"

}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go Database1(ch1)
	go Database2(ch2)
	merge(ch1, ch2)

}

// we have slice of channels (varidaic parameter)
func merge(chs ...chan string) {

	for _, ch := range chs { // iterating over slice of channels

		fmt.Println(<-ch) //read data from a channel
	}
}
