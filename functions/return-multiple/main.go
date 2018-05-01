package main

import "fmt"

func main() {
	var fName string = "Dr."
	var lName string = "Strange"

	fmt.Println(greet(fName, lName))
}

func greet(fName string, lName string) (string, string) {
	return fmt.Sprint(fName, lName), fmt.Sprint(fName, lName)
}
