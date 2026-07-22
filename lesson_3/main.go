package main

import "fmt"

func main() {
	// variable of type uint8
	var smallPositiveValue uint8
	smallPositiveValue = 10
	fmt.Println(smallPositiveValue)
	// variable of type int8
	var smallPositiveNegativeInt int8
	smallPositiveNegativeInt = -10
	fmt.Println(smallPositiveNegativeInt)
	// uint16, uint32, uint64
	// int16, int32, int64
	var myInt int = 23000000000000000
	fmt.Println(myInt)
	myInt = int(smallPositiveNegativeInt)
	myInt = int(smallPositiveValue)
	// to typecase a value int()

	var myByte byte = 'A'
	fmt.Println(myByte)
	// byte is mainly used to represent raw data
	// and also to distinguish between uint8 and byte
	// since byte is an alias dor uint8

	// rune is an alias for int32
	var myRune rune = 'B'
	fmt.Println(myRune)
	myRune = '🗿'
	fmt.Println(myRune)
}

