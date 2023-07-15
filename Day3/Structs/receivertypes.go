package main

import "fmt"

type Student1 struct {
	StdName string
	StdId   int
}

//Value type receiver
// func (std Student) setSTd() {
// 	std.StdId = 400

// }

//Ref Type Receiver

func (std *Student1) setSTd() {
	std.StdId = 400

}

func main() {

	std1 := Student1{"Sachin", 100}

	std1.setSTd()

	fmt.Println(std1)

}
