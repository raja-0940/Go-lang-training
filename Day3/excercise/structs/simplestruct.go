package main

import "fmt"

type Employee struct {
	//value type struct
	First_Name string
	Last_Name  string
	ID         int
}

func main() {
	//way1 to create object with properties of struct
	var emp1 = Employee{First_Name: "Raja", Last_Name: "Ram", ID: 100}
	fmt.Println(emp1.First_Name, emp1.Last_Name, emp1.ID)
	//way2 to create a object without properties
	var emp2 =  Employee{"Rama","Raju",1000}
	fmt.Println(emp2.First_Name, emp2.Last_Name, emp2.ID)
	//way3 to create a nil object and add properties
	var emp3 Employee
	fmt.Println(emp3.First_Name, emp3.Last_Name, emp3.ID)
	emp3.First_Name = "Roja"
	emp3.Last_Name = "Rani"
	emp3.ID = 190
	fmt.Println(emp3.First_Name, emp3.Last_Name, emp3.ID)


}