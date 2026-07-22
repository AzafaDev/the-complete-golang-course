package main

import (
	"fmt"
)

func main() {
	var smallFloat float32
	fmt.Println(smallFloat)
	smallFloat = 23.0234324
	fmt.Println(smallFloat)

	var bigFloat float64
	fmt.Println(bigFloat)
	bigFloat = 43.01029021901293
	fmt.Println(bigFloat)
	bigFloat = 43.01029021901293121820120
	fmt.Println(bigFloat)

	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)

	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPart)
}
