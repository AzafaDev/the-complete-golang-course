package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 202
	var b uint8 = 141

	// the or operator
	// c := a | b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// the and operator
	// c := a & b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// the xor operator
	// c := a ^ b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// the and not or bit clear operator
	// c := a &^ b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// the invert operator
	c := ^a
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))
}
