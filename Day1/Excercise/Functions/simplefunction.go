package main

import "fmt"

func main(){
	//calling simple function 
	func_demo("Rajakumar  Battula",37731)
	//calling function with one return type 
	value := func_demo2("Rajakumar Battula")
	fmt.Println(value)
	//calling function with more than one return types
	Name, Salary := func_demo3(37731)
	fmt.Println(Name, Salary)
}

//simple function
func func_demo(Name string, ID int){
	fmt.Println("My name is : ",Name)
	fmt.Println("My ID number is : ",ID)
}

//function with return one value
func func_demo2(Name string) int {
	return 37731
}

//function with returning more than one value
func func_demo3(ID int) (EmpName string, EmpSalary int) {
	// Name := "Rajakumar"
	// Salary := 1200000
	EmpName="Rajakumar Battula"
	EmpSalary=1200000
	// return Name, Salary
	return
}