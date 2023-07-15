package main

import "fmt"

type Item struct {
	ItemName string
	ItemId int
}

//Method: function with receiver type (struct type)
//GetItem is a method of Item struct
func (item Item) GetItem() {
	fmt.Println("Get Item is called", item.ItemName, item.ItemId)
}

func (item Item) SetItem() {
	fmt.Println("Set Item is called", item.ItemName, item.ItemId)
}

// func GetItem(){  //function

// }

func main() {

	// GetItem()

	item1 := Item{"Laptop", 10}

	item1.GetItem()
}
