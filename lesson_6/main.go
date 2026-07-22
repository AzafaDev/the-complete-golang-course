package main

import "fmt"

// Group similar variables
var (
	myInt    int    = 10
	myString string = "Hello World!"
)

func main() {
	fmt.Println(myInt)
	fmt.Println(myString)

	// implicit type assignment
	var age = 25
	fmt.Println(age)

	// multiple variable declaration
	var year, name = 2023, "code & learn"
	fmt.Println(year)
	fmt.Println(name)

	// shorthand variable declaration
	newAge := 25
	newYear := 2023
	fmt.Println(newAge, newYear)

	// // examples:
	// var apartmentNumber int
	// var apartmentNumber int = 2000
	// var apartmentNumber = 2000
	// var apartmentNumber, streetName = 2000, "Bahshore"
	apartmentNumber := 2000

	fmt.Println(apartmentNumber)
	

}
