package main

import "fmt"

func main() {

	//array2 := [3]string{"java", "C", "C+"}  // array

	//Slice : array without any size

	slice1 := []string{"java", "C", "C++"} //slice

	//we can add elements to the slice

	slice1 = append(slice1, "GoLang", "Python")

	for _, v := range slice1 {

		fmt.Println(v)
	}
}
