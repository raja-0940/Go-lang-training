package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ch1 := make(chan int)

	wg.Add(2)
	go WriteDB(ch1)
	go ReadFile(ch1)

	fmt.Println("Control returned to main..")
	fmt.Println("Main will execute it's tasks...")
	wg.Wait()
	//We have to make main wait till both go routines finish their execution
	//time.Sleep(5 * time.Second)
}

func ReadFile(ch1 chan int) {
	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("Put value into channel", i)
		//We have to put value/record into channel

		ch1 <- i //putting value into channel
	}

	wg.Done()
}

func WriteDB(ch1 chan int) {

	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("Read value from channel ", <-ch1) //reading data from channel
	}

	wg.Done()
}
