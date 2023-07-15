// Go program to illustrate how
// to sort a slice stable
package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {

	// Creating and initializing a structure
	Author := []struct {
		a_name    string
		a_article int
		a_id      int
	}{
		{"Mina", 304, 1098},
		{"Cina", 634, 102},
		{"Tina", 104, 105},
		{"Mina", 34, 109},
		{"Cina", 634, 102},
		{"Mina", 4, 100},
		{"Rohit", 56, 1098},
		{"Cina", 634, 102},
		{"Mina", 39, 1098},
		{"Sohit", 20, 110},
	}

	// Sorting Author by their names
	// Using SliceStable() function
	sort.SliceStable(Author, func(p, q int) bool {
		return Author[p].a_name < Author[q].a_name
	})

	fmt.Println("Sort Author according to their names:")
	fmt.Println(Author)

	// Sorting Author by the total articles
	// Using SliceStable() function
	sort.SliceStable(Author, func(p, q int) bool {
		return Author[p].a_article < Author[q].a_article
	})

	fmt.Println()
	fmt.Println("Sort Author according to their" +
		" total number of articles:")

	fmt.Println(Author)

	// Sorting Author by their ids
	// Using SliceStable() function
	sort.SliceStable(Author, func(p, q int) bool {
		return Author[p].a_id < Author[q].a_id
	})

	fmt.Println()
	fmt.Println("Sort Author according to the their Ids:")
	fmt.Println(Author)
}
