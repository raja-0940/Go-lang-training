package main

import "fmt"

type Product struct {
	ProductName string
	ProductID int

}

func main() {
	prd1 := Product{"TV",100}
	changeProduct(&prd1)
	fmt.Println(prd1)

	prd2 := &Product{"Refridgerator", 20} //creating ref type object
	changeProduct(prd2)
	fmt.Println(prd2)

	prd3 := new(Product)  //pointer type nil object will be created
	prd3.ProductName = "Car"
	prd3.ProductID = 30
	changeProduct(prd3)
	fmt.Println(prd3)

}

func changeProduct(prd *Product){
	prd.ProductName = "Modify"
}