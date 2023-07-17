package main

import "fmt"

func main() {
	//array defination
	arr := [2]string{"Aishwarya", "Raj"}
	fmt.Println(arr)
	//slice defination
	sl1 := []string{"Raja","kumar"}
	fmt.Println(sl1)
	//we can add elements to the slice
	sl1=append(sl1,"swapna","Aishwarya")
	for _,v := range sl1 {
		fmt.Println(v)
	}
}