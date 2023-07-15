package main

import "fmt"

//var counter = 0 //global variable

//counter should be local
//still it should be initialized to 0 only once

// func getHits() int {

// 	var counter=0
// 	counter++

// 	return counter
// }

// var getHits,getHits1 = func()  (func() int, func() int ) { // this outerfunc is returning defn of addHits

// 	var counter = 0 // its local variable
// 	//List of EMployee

// 	//add closure functions which will have crud perations for this List

// 	var addHits = func() int {

// 		counter++
// 		return counter
// 	}

// 	var decHits = func () int {

// 		return counter
// 	}

// 	//return addHits ,decHits// we r returning inner function/closure function

// 	return addHits, decHits

// }()// it will execute outer function & will return defn addHits

var getHits = func() func() int { // this outerfunc is returning defn of addHits

	var counter = 0 // its local variable
	//List of EMployee

	//add closure functions which will have crud perations for this List

	var addHits = func() int {

		counter++
		return counter
	}

	//return addHits ,decHits// we r returning inner function/closure function

	return addHits

}()

func main() {

	// fmt.Println(temp())

	// fmt.Println(temp())

	// fmt.Println(temp())

	fmt.Println(getHits()) //we will call addHits

	fmt.Println(getHits()) //we will call addHits
	fmt.Println(getHits()) //we will call addHits

	//  var exp=  getHits() //it will return defn of addHits

	// fmt.Println(getHits()())
	// fmt.Println(getHits()())
	// fmt.Println(getHits()())

	// fmt.Println(getHits())
	// fmt.Println(getHits())
	// //counter = 100
	// fmt.Println(getHits())

}
