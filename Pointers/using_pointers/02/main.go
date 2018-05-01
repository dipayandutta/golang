package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p \n", &z) // address in function zero
	fmt.Println(&z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p \n", &x) // address in function main
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}
