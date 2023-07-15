package main

// "rsc.io/quote/v4"

import (
	"fmt"

	_ "myproject.com/DAL"   //init will be called even though we import it using _
	b "myproject.com/DataAccess/FileLog"
	a "myproject.com/Log/FileLog"
)

//init is called by runtime before calling main
//it will be called only once
//we can't call it explicitly

func init() {

	fmt.Println("init of main is called.")
}

func main() {

	// fmt.Println(quote.Hello())

	// fmt.Println(stringutil.Reverse("ABC"))

	// color.Blue("Prints %s in blue.", "text")

	//DAL.ReadDB()

	a.WriteLog()
	b.FIleLogDA()
	//DAL.WriteDB()

	//DAL.ReadMySQL()
}
