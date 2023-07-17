package main

import "fmt"

func main() {

	fmt.Println("main has started...")
	OpenFile()
	fmt.Println("returned to main...")
}

func OpenFile() {

	defer CLeanUp() //defer :finally
	fmt.Println("File is opened..")
	//panic("Exception thrown..") //throwing exception

	//new Exception()

	var no1, no2 = 10, 0

	result := no1 / no2

	fmt.Println("Open file finished.", result)

}
func CLeanUp() {
	fmt.Println("Resource Clean up")

	err := recover() // catching

	fmt.Println(err)
}
