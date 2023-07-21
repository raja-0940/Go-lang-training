package main

import "fmt"

func main() {
	var sl = []int{1,2}
	sl1 := [...]int{1,2,3}

	fmt.Println(sl,"\n",sl1)

	fmt.Printf("Type of slice sl1 is %T", sl1[1])
	fmt.Println(sl[0])

	for i := 0; i<len(sl1); i++ {
		fmt.Println(i)
	}

	//array-->slice
	arr1 := [6]int{10, 11, 12, 13, 14,15}
	myslice := arr1[0:6] //arr1[start_index:end_index]
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	//slice from make()
	myslice1 := make([]int,2,4)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	//insert values at paricular index of slice
	sl3 := []int{1:10,2:40}
	fmt.Println(sl3)

	//change elements from the slice
	sl4 := []int{10,20,30}
	fmt.Println(sl4)
	sl4[1] = 90
	fmt.Println(sl4)

	//append elements to the slice
	sl4 = append(sl4,40,50,60)
	fmt.Println(sl4)

	//append one slice to the another slice
	// var sl5 = []int{}
	sl5 := append(sl4, sl3...)
	fmt.Println(sl5)
	sl51 := sl5[:len(sl5)-5] //slicing slice
	sl6 := make([]int, len(sl51)) //creating empty slice
	copy(sl6,sl51) //coping slice
	fmt.Println(sl6)

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Println(neededNumbers)
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}