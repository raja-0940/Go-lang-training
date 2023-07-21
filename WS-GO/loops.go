package main

import (
	"fmt"
)

func main() {

	//simple for loop
	x := 3
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
	//continue statement : to skip one or more iterations (following example skips 3rd iteration)
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i) // 0 1 2 4
	}
	//break : used to breaks out of the loop when it equals 3
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i) // 0 1 2 4
	}
	//nested for loops
	behaviour := [2]string{"Tasty", "Sweet"}
	fruit := [3]string{"Banana", "Apple", "Orange"}
	for i := 0; i < len(behaviour); i++ {
		for j := 0; j < len(fruit); j++ {
			fmt.Println(behaviour[i],fruit[j])
		}
	}
	//for loop using range keyword
	for index, value := range fruit {
		fmt.Println(index, value)
	}
	for _, value := range fruit {  // '_' is blank operator which is used in case if we not required index or value
		fmt.Println(value)
	}
	for index, _ := range fruit {
		fmt.Println(index)
	}
	

}