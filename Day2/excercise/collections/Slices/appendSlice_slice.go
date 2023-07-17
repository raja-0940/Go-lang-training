package main
import (
	"fmt"
)
func main() {
	slice1 := []int{10,20,30}
	slice2 := []int{40,50,60}
	slice1=append(slice1,slice2...)
	fmt.Println(slice1)
}