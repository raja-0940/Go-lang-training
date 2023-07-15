package main

import (
	"fmt"
	"sort"
)

type employee struct {
	name   string
	salary int
}

type employeeList []employee

func (e employeeList) Len() int {
	return len(e)
}

func (e employeeList) Less(i, j int) bool {
	return e[i].salary > e[j].salary
}

func (e employeeList) Swap(i, j int) {
	fmt.Println(e)
	e[i], e[j] = e[j], e[i]
}

func main() {

	array1 := [4]string{"C", "java", "Go", "Node"}

	sort.Strings(array1[:])

	fmt.Println(array1)

	eList := []employee{
		employee{name: "John", salary: 3000},
		employee{name: "Bill", salary: 4000},
		employee{name: "Sam", salary: 1000},
	}
	sort.Sort(employeeList(eList))
	for _, employee := range eList {
		fmt.Printf("Name: %s Salary %d\n", employee.name, employee.salary)
	}
}
