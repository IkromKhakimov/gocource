package basics

import "fmt"

var middleName = "Cane"

func main() {
	//var age int
	//var name string = "John"
	//var namel = "Jane"
	//
	//count := 10
	//lastName := "Smith"

	// Default values
	// Numeric Types: 0
	// Boolean Types: false
	// String Types: ""
	// Pointers, slices, maps, function, and structs: null

	middleName := "Marie"
	fmt.Println(middleName)

	// ---- SCOPE
	// fmt.Println(firstName)
}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}
