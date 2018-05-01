package main

import "fmt"

func main() {
	name := "Dr. Strange"

	switch {
	case len(name) == 8:
		fmt.Println(name)
	case name == "Hello":
		fmt.Println("Hello")
	case name == name:
		fmt.Println("Name is ", name)
	default:
		fmt.Println("Default")
	}
}
