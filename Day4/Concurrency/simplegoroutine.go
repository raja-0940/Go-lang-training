package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go WriteFile()
	go ReadFile()

	fmt.Println("Control returned to main..")
	fmt.Println("Main will execute it's tasks...")

	wg.Wait()

	//We have to make main wait till both go routines finish their execution

	//time.Sleep(5 * time.Second)
}

func WriteFile() {

	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("value of i from WriteFIle", i)
	}

	//wg.Done()
}

func ReadFile() {
	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("value of i from ReadFile", i)
	}

	wg.Done()
}
