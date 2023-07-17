package main

import "fmt"

func main() {
	// 'defer' keyword is like 'finally' in other programming languages.
	defer foo()
	defer test() // Last In First Out
	bar()
	fmt.Println("Doing something ... ")
}

func foo(){
	fmt.Println("foo is called")
}

func bar(){
	fmt.Println("bar is called")
}

func test(){
	fmt.Println("Test is called")
}