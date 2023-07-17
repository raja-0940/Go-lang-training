package main

import (
	"fmt"

	"time"
)

func main() {

	ch1 := make(chan int)
	go CheckoutProducts(ch1)

	fmt.Println("Control returned to main...")
	BrowseProducts()

	//we have to wait till CheckoutProduct is finished

	var finalValue = <-ch1 //blocking operation. It will wait till data is read from channel

	StartDelivery(finalValue)
}

func CheckoutProducts(ch1 chan int) {

	time.Sleep(4. * time.Second)

	fmt.Println("Checkut Product is finished...")

	ch1 <- 9999
}

func BrowseProducts() {

	fmt.Println("Browse Products..")
}

func StartDelivery(finalValue int) {

	fmt.Println("Start Delivery", finalValue)
}
