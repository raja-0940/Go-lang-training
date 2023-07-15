package main

import "fmt"

// var count=0

// func getHit() int {
// 	count++
// 	return count
// }

var gitHit = func() func() int{
	var count=0
	var addHit = func() int{
		count++
		return count
	}
	return addHit
}()

func main() {
	// fmt.Println(getHit())
	// fmt.Println(getHit())
	// fmt.Println(getHit())
	fmt.Println(gitHit())
	fmt.Println(gitHit())
	fmt.Println(gitHit())
}