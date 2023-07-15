package main

import "fmt"

func main() {

	func1("Mukta", 1000)
	value := func2("Anvesh")
	fmt.Println(value)
	fname, salaray := GetEMpDetails(10)
	fmt.Println(fname, salaray)

}



//We get 0 intialized return parameters
//We can have just return
//While calling we get a kind of documentation

func GetEMpDetails(Id int) (EmpName string, EmpSalary int) {

	// var name = "Chetan"

	// var salary = 1000

	EmpName ="Chetan"
	EmpSalary=1000

	return 

}

func func1(Name string, Id int) {

	fmt.Println(Name, Id)
}

func func2(Name string) int {

	return 1000

}

//function returning more than one value

// func GetEMpDetails(Id int) (string, int) {

// 	var name = "Chetan"

// 	var salary = 1000

// 	return name, salary

// }
