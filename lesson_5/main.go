package main

import "fmt"

func main() {
	var myString string
	fmt.Println(myString)
	myString = "Hello\n World"
	fmt.Println(myString)

	myString = `Hello
	
	World`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "code"
	lastName = "learn"

	// var fullName string
	// fullName = firstName + " " + lastName
	// fmt.Println(fullName)

	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)

	

}
