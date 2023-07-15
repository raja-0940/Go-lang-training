package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// package main
// import "fmt"
// func main() {
// 	mystr := "Hello, alexa!"  //create string
// 	fmt.Println("The original string given here is:", mystr)
// 	runes := []rune(mystr)  //turn string to slice
// 	fmt.Println("The iteration performed through each character of string:")

// 	for _, v := range runes { //iterate through rune

// fmt.Printf("%c ", v)  //print characters along with space character
// 	}
// }

// package main
// import "fmt"

// func main() {
// 	mystr := "Hello, alexa!" //create string
// 	fmt.Println("The original string given here is:", mystr)
// 	fmt.Println("The iteration performed through each character of string:")
// 	for i := 0; i < len(mystr); i++ {   //run a loop and iterate through each character
// 		fmt.Printf("%c ", mystr[i])  //print characters along with the space character
// 	}
// }
