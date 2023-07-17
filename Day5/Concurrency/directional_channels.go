package main

import "fmt"

// Main function
func main() {

	// // Only for receiving
	//mychanl1 := make(<-chan string,2)

	// // Only for sending
	//mychanl2 := make(chan<- string)

	// // Display the types of channels
	// fmt.Printf("%T", mychanl1)
	// fmt.Printf("\n%T", mychanl2)

	//Convert Bidirectional channel into unidirectional

	ch1 := make(chan string)

	go Routine(ch1)
	ch1 <- "Put Value"
}

func Routine(ch1 <-chan string) {
	fmt.Println(<-ch1)
}

//https://www.geeksforgeeks.org/unidirectional-channel-in-golang/



// / Go program to illustrate how to convert
// // bidirectional channel into the
// // unidirectional channel
// package main
  
// import "fmt"
  
// func sending(s chan<- string) {
//     s <- "GeeksforGeeks"
// }
  
// func main() {
  
//     // Creating a bidirectional channel
//     mychanl := make(chan string)
  
//     // Here, sending() function convert
//     // the bidirectional channel
//     // into send only channel
//     go sending(mychanl)
  
//     // Here, the channel is sent 
//     // only inside the goroutine
//     // outside the goroutine the 
//     // channel is bidirectional
//     // So, it print GeeksforGeeks
//     fmt.Println(<-mychanl)
// }
