package main

import "fmt"

func main(){
	calMarks("Rajakumar")
	calMarks("Swapna",60,70,80)
	calMarks("Aishwarya",100,99,100)
}

//Varidic parameter : to this parameter we can pass 0 or more parameters
//Internally it stores in a slice (dynamic array)
func calMarks(Name string, marks ...int){
	fmt.Println(Name)
	fmt.Println(marks)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
}