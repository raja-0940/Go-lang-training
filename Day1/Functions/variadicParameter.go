package main

import "fmt"

func main() {

	CalculateMarks("Pallav")
	CalculateMarks("Abhinav", 70, 80, 90)
	CalculateMarks("Anvesh", 70, 80, 90, 100)

	Test("Chandra","Neelesh", "Avinash", 100, tru e)

}

//Variadic parameter: to this parameter we can pass 0 or more values

//Internally it stored in slice (Dynamic array)

func CalculateMarks(Name string, marks ...int) {

	fmt.Println(Name)
	fmt.Println(marks)
	fmt.Println("++++++++++++++++")

}

func Test(Name string, data ...any) {

	fmt.Println(Name)
	fmt.Println(data)
}
