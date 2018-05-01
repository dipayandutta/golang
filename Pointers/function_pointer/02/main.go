package main

import "fmt"

func zero(z *int, y *int) {

	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(*z)
	fmt.Println(*y)

	d := *z + *y
	fmt.Println(d)
	a := *z
	b := *y
	fmt.Println(a)
	fmt.Println(b)

	c := a + b
	fmt.Println(c)
}

func main() {
	z := 10
	b := 20
	zero(&z, &b)
}
