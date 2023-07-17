package main
import (
	"fmt"
)
func main() {
	WEEEKDAYS := []string{"MON","TUE","WED","THU","FRI"}
	
	fmt.Println(WEEEKDAYS[2:3])  // prints only WED

	WEEEKDAYS=append(WEEEKDAYS[0:2], WEEEKDAYS[3:]...)

	fmt.Println(WEEEKDAYS)
	

}