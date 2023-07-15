package main

import "fmt"

func main() {

	weekdays := []string{"MON", "TUE", "WED", "THR", "FRI"}

	weekdays = append(weekdays[:2], weekdays[3:]...)

	fmt.Println(weekdays)
	//fmt.Println(weekdays1)

}
