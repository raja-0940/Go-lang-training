package main

import "fmt"

func main() {
	var arr[3]int
	fmt.Println(arr) // prints [0 0 0]
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	for i:=0; i<len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i,v := range arr {
		fmt.Println(i ,"and", v)
	}

	for _,v := range arr {
		fmt.Println(v)
	}
	for i,_ := range arr {
		fmt.Println(i)
	}

	//Another way to define fixed length array
	var arr1=[3]string{"Java","Python","Go"}
	fmt.Println(arr1)
	
}