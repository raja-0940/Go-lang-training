package main

import "fmt"

func main() {

	var array1 [3]int //this array will be of type int

	array1[0] = 100
	array1[1] = 200
	array1[2] = 300
	//fmt.Println(array1)

	for i := 0; i < len(array1); i++ {

		fmt.Println(array1[i])
	}

	// for i,v := range array1 {

	// }

	for _, v := range array1 {

		fmt.Println(v)
	}

	array2 := [3]string{"java", "C", "C+"} //this array will be of type string

	fmt.Println(array2)
}
