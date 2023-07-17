package main
import (
	"fmt"
)
func main() {
	slice1 := []int{10,20,30,40}
	
	fmt.Println(slice1[0:3])  // elements of slice will move from L to H-1
	fmt.Println(slice1[:2])  // 0 to H-1
	fmt.Println(slice1[1:])  // 1 to H
	

}