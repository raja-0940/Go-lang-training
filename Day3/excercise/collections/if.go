package main

import "fmt"

func main(){

	// val := 100
	// if val==100 {
	// 	fmt.Println("exists")
	// }else{
	// 	fmt.Println("not exists")
	// }
	if value := test(); value == 100{
		fmt.Println("if is called")
	}else{
		fmt.Println("else is called")
	}

}

func test()int{
	return 100
}