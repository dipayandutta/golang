package main

import "fmt"

func main() {
	fName := "Dr."
	lName := "Strange"
	fmt.Println(greet(fName, lName))
}

func greet(fName string, lName string) (s string) {
	s = fmt.Sprint(fName, lName)
	return
}
