package main

import "fmt"

func main() {
	var myString string
	// string zero value is ""
	fmt.Println(myString)
	myString = "Welcome to code and learn"
	fmt.Println(myString)

	var myInt int
	// zero value for number types is 0
	fmt.Println(myInt)
	myInt = 10
	fmt.Println(myInt)

	var myBool bool
	// zero value for bool is false
	fmt.Println(myBool)
	myBool = true
	fmt.Println(myBool)

	myString = "string new value"
}
