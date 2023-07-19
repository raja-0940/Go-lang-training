package main

import "fmt"

// func Calculate(marks...int){
// }

//map[int][]Student

//slice : []Student

type Student struct {
	FName string
	LName string
}

type Exams struct {
	UnitTestsExam int
}
func (exam Exams) TakeExam() {
	fmt.Println("TakeExams from Exams struct")
}
func (std Student) TakeExam() {
	fmt.Println("TakeExams from Student struct")
}
type SchoolStd struct { // child struct
	Student
	Exams
}

func main() {
	std := Student{"Sachin", "T"}
	exam := Exams{10}
	schoolstd := SchoolStd{std, exam}
	fmt.Println(schoolstd.FName, schoolstd.UnitTestsExam)
	//schoolstd.TakeExam()  //: ambiguty
	schoolstd.Student.TakeExam()
	schoolstd.Exams.TakeExam()
}

