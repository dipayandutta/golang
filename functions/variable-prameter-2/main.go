package main

import "fmt"

func main() {
	n := function(10, 21, 3, 4)
	fmt.Println(n)
}

func function(numbers ...int) int {
	fmt.Println(numbers)
	return 1
}
