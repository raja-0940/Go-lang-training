package main

import (
	"fmt"
)

type emp struct {
	empName string
	empId int	
}

func empDetails( e1 emp ) {
	fmt.Println("Name : ", e1.empName)
	fmt.Println("ID : ", e1.empId)
}

func main() {
	var e emp
	e.empName = "Rajakumar Battula"
	e.empId = 39
	// fmt.Println("Name : ", e.empName)
	// fmt.Println("ID : ", e.empId)
	empDetails(e)
}