package main

import "fmt"

type Student struct {
	StudentName string
	StudentID int
}

//Value type receiver
// func (std Student) SetSTd(){
// 	std.StudentName = "Ramakanth"
// 	std.StudentID = 20
// }

//Ref type receiver
func (std *Student) SetSTd(){
	std.StudentName = "Ramakanth"
	std.StudentID = 20
}

func main() {
	std1 := Student{StudentName: "Mallikarjun", StudentID: 39}
	std1.SetSTd()
	fmt.Println(std1)
}