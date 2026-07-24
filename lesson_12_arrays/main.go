package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a = [5]int{5, 2: 10, 50, 11}
	fmt.Println(a)

	b := [...]int{3, 4, 1, 2}

	fmt.Println(b)
	fmt.Println(len(b))

	c := a[3]
	fmt.Println(c)

	d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(d)

}
