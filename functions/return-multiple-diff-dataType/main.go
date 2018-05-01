package main

import "fmt"

func main() {

	var age int = 40
	var fName string = "IronMan"
	fmt.Println(greet(age, fName))
}

func greet(age int, fName string) (int, string) {
	return (age), fmt.Sprint(fName)
}
