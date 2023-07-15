package main

import "fmt"

type Log interface {

	//IDimen

	WriteLog()

	//Perimeter()
}

type DAL interface {
	WriteData()
}

type Traingle struct {
	side int
}

func (tra Traingle) WriteLog() {

	fmt.Println("WriteLog of Traingle is called.")

}

func (tra Traingle) WriteData() {

	fmt.Println("WriteData of Traingle is called.")

}

// func (tra Traingle) Topology() {

// 	fmt.Println("Topology of Traingle is called.")

// }

type Square struct {
	side int
}

func Logging(log Log) {

	log.WriteLog()
}

func Data(data DAL) {

	data.WriteData()
}
func main() {

	tra := Traingle{40}
	Logging(tra)
	Data(tra)

}
