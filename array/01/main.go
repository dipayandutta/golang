package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[1] = 10
	fmt.Println(x[1])
}
