package main

import "fmt"

func main() {
	foo()
	test()
	bar()

	var i = 0

	var j = 10 / i

	defer fmt.Println(j)
	fmt.Println("Doing something...")
}

func foo() {
	fmt.Println("foo is called")
}

func bar() {
	fmt.Println("bar is called")

}

func test() {
	fmt.Println("test is called")

}
