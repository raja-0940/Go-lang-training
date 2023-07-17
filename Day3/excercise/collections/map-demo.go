package main

import "fmt"

func main() {


	map1 := map[int]string{1: "java", 2: "python", 3: "shell scripting", 4: "Go language"}

	fmt.Println(map1[4])

	for k,v := range map1 {
		fmt.Println(k,":",v)
	}

	var map2 map[int]string
	fmt.Println(map2) //  empty map or nil map

	var map3 = make(map[int]string)
	map3[1]="Raja"
	map3[2]="Aishwarya"
	map3[3]="Swapna"
	map3[4]="Outer"
	fmt.Println(map3)
	map3[1]="Rajakumar Battula"
	fmt.Println(map3)
	//delete from map
	delete(map3,4)
	fmt.Println(map3)	
	//check key if exists or not
	if value, exist := map3[3]; exist {
		fmt.Println(value, exist)
	}else{
		fmt.Println(value, exist)
	}

	//map inside a map
	map4 := make(map[string]map[string]int)
	map4["Golang"]=map[string]int{"Day1": 3, "Day2": 3}
	fmt.Println(map4["Golang"])

	//nested map with slices
	slice1 := []int{10,20,30}
	slice2 := []int{40,50,60}
	map5 := make(map[string]map[string][]int)
	map5["Go lang"] =  map[string][]int{ "Day1": slice1, "Day2": slice2}
	fmt.Println(map5["Go lang"])
	for _,value := range map5 {		
		for k,v := range value{
			fmt.Println(k,v)
			for _,v1 := range v{
				fmt.Println(v1)
			}
		}
	}
}