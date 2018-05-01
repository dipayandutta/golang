package main

import "fmt"

func arrayOne(args ...string) {
	fmt.Println(args)
}
func arrayTwo(args []string) {
	fmt.Println(args)
}
func stringPointer(args *string) {
	fmt.Println(*args)
}
func main() {

	arrayOne([]string{"Alphabet", "Google"}...)
	arrayTwo([]string{"A", "B", "C"})

	str := "Hello world!"
	stringPointer(&str)
}
