package main

import "fmt"

func main() {

	var slice1 = []int{10, 20, 30, 40}

	var slice2 = []int{50, 60, 70, 80}

	Test(10, 20, 30, 40)
	Test(slice1...)

	//append(slice1,10,20,30)

	slice1 = append(slice1, slice2...)

	fmt.Println(slice1)

}

func Test(values ...int) {

}
