package main

import "fmt"

type Product struct {
	ProductName string
	ProductId   int
}

func main() {

	prd1 := &Product{"TV", 100} //creating ref type of struct obj

	changeProduct(prd1)

	fmt.Println(prd1)

	prd2 := new(Product) //creating ref type of struct obj

	fmt.Println(prd2)
	prd2.ProductName = "Laptop"
	prd2.ProductId = 900

	// changeProduct(&prd1)

	// fmt.Println(prd1)

}

func changeProduct(prd *Product) {

	prd.ProductName = "Modify"
}
