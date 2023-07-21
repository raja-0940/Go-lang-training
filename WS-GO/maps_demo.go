package main

import (
	"fmt"
)

func main() {
	empInfo := map[int]string{1: "Rajakumar Battula", 2: "Swapna S", 3: "Aishwarya Raj Battula"}
	fmt.Println(empInfo)
	for k,v := range empInfo {
		fmt.Println(k, "is", v)
	}
	//maps using make()
	map1 := make(map[int]string)  // empty map
	map1[1] = "Rama"
	map1[2] = "Sita"
	map1[3] = "Lakshmana"
	map1[4] = "Hanuma"
	map1[5] = "Ravana"	
	fmt.Println(map1)
	//deleting elements from the map
	delete(map1,5)
	//checking for existing and its value
	value, ok := map1[1]
	val, _ := map1[4]
	fmt.Println(val)
	fmt.Println(value,ok)
	fmt.Println(map1)
	//empty map
	var map2 map[int]int
	fmt.Println(map2)
	//maps are references to hash tables
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)

	//As maps are unordered, to sort maps in order
	a1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Println(a1)

	var b1 = []string{}   // defining the order
	b1 = append(b1, "one", "two", "three", "four")

	for k, v := range a1 {        // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b1 {  // loop with the defined order
		fmt.Printf("%v : %v, ", element, a1[element])
	}
}