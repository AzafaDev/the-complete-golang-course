package main

import "fmt"

const (
	x                     = 10
	y               int32 = 15
	applicationName       = "lesson 7"
	isRunning             = false
	character = 'a'
	isTrue = 1 < 2
)

func main() {
	fmt.Println(isTrue)
	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = x
	fmt.Println(b)

	var c int
	c = int(y)
	fmt.Println(c)

	var d int32
	d = y
	fmt.Println(d)

	const z = complex(10.2, 100.9)
	const l = imag(z)
}

func sum(a, b int) int {
	return a + b
}
