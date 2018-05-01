package main

import "fmt"

func function(name string) {
	fmt.Println(name)
}
func pointer(z *int) {
	fmt.Println("Address is ", z)
	fmt.Println(*z)
}
func main() {
	var name string = "Dr. Strange"
	function(name)
	name = "Iron Man"
	function(name)

	z := 10
	fmt.Println("Address inside main function ", &z)
	pointer(&z)
}
