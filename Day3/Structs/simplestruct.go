package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Id        int
}

func main() {

	//Way1
	var emp1 = Employee{FirstName: "Sachin", LastName: "T", Id: 100}

	fmt.Println(emp1.FirstName, emp1.LastName, emp1.Id)

	//Way2 (Without Property Names)

	var emp2 = Employee{"Virat", "K", 300}
	fmt.Println(emp2.FirstName, emp2.LastName, emp2.Id)

	//way 3

	var emp3 Employee

	emp3.FirstName = "Mukta"

	fmt.Println(emp3.FirstName, emp3.LastName, emp3.Id)

}
