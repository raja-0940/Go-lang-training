package main

import "fmt"

//map : key value data structure

// access elements : by specifying key

func main() {

	// var map4 = make(map[string]map[string]int)

	// map4["GoLang"] = map[string]int{"Day1":2,"Day2":3}
	// map4["React.js"] = map[string]int{"Day1":4,"Day2":5}




	var map4 = make(map[string]map[string][]int)


	slice1 := []int{10,20,30}
	slice2 := []int{40,50,60}


	map4["GoLang"] = map[string][]int {"Day1GoLang":slice1,"Day2GoLang":slice2}


// for key,value := range map4{

// 	for k1,v1 := range value {


// 	}
// }



	// map4["GoLang"] = map[string]int{"Day1":2,"Day2":3}
	// map4["React.js"] = map[string]int{"Day1":4,"Day2":5}

  //"GoLang" : Day1:3,Day2:4



	var map1 = map[string]int{"Java": 10, "Node.js": 20, "React.js": 40, "GoLang": 50} // key data type: string: value type:int
	map1["Angular"] = 0
	fmt.Println(map1["GoLang"]) //accessing value using key
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	//var map2 map[int]string //nil map
	var map2 = make(map[int]string) // key type: int, value type:string
	map2[100] = "John"
	map2[200] = "ken"
	map2[300] = "Ben"
	fmt.Println(map2)

	map2[300] = "Arpan" // it will update

	fmt.Println(map2)

	//delete from map

	delete(map2, 300)

	fmt.Println(map2)

	//check key exists or not

	if value, exist := map2[200]; exist {

		fmt.Println(value, exist)
	} else {
		fmt.Println(value, exist)
	}

}
