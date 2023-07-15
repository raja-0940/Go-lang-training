package DAL

import "fmt"

func init() {

	fmt.Println("init of DAL is called.")
}

func readMySQL() {

	fmt.Println("ReadMYSQL is called..")

	//updateDB()
}

type employee struct {
	EmpName string
	empID   int
}
