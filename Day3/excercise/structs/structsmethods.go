package main

import "fmt"

type Item struct {
	ItemName string
	ItemID int
}

//Method : function with receiver type (struct type)
//GetItem is a method of Item struct
// to call a method, we need use its object
func (item Item) GetItem(){
	fmt.Println("Calling GetItem method")
	fmt.Println(item.ItemName, item.ItemID)
}

func (item Item) SetItem(){
	fmt.Println("Calling SetItem method")
	fmt.Println(item.ItemName, item.ItemID)
}

// func GetItem() {

// } this is a function. it can call independently

func main() {
	item1 := Item{"Bike", 10}
	item1.GetItem()
	item1.SetItem()

}