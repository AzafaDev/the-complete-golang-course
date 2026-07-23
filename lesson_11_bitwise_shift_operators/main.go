package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 20
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = a >> 2
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = a << 2
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = 1
	a = a << 3
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = 5 | (1 << 1)
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = a &^ (1 << 1)
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = a ^ (1 << 0)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
}
