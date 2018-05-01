package main

import "fmt"

func printName(fName *string, lName *string) {
	fmt.Println(*fName)
	fmt.Println(*lName)
}
func main() {
	var fName string = "Dr."
	var lName string = "Strange"

	printName(&fName, &lName)

}
