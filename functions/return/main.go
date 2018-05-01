package main

import "fmt"

func main() {
	fmt.Println(greet("Iron Name", "Dr. Strange"))
}

func greet(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}
