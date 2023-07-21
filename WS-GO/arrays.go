package main

import "fmt"

func main() {
	var arr = [2]int{1,2}
	arr1 := [...]int{1,2,3}

	fmt.Println(arr,"\n",arr1)

	fmt.Println(arr1[1])
	fmt.Println(arr[0])

	for i := 0; i<len(arr1); i++ {
		fmt.Println(i)
	}

	arr2 := [3]string{}
	arr2[0] = "Ram"
	arr2[1] = "Sita"
	fmt.Println(arr2)
	arr2[2] = "Hanuman"
	fmt.Println(arr2)

	arr3 := [5]int{1:10,2:40}
	fmt.Println(arr3)
}