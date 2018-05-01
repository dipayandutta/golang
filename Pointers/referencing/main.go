package main

import "fmt"

func main() {
	a := 10

	fmt.Println("Value of A --> ", a)
	fmt.Println("Address of A --> ", &a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 41
	fmt.Println(*b)

}
