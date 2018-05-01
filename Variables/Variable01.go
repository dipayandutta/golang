package main

import "fmt"

func main() {

	a := 10
	b := 12.23
	c := true
	d := "GO Lang"
	e := "M"

	//Print the Values of a Variable
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n ", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	//While using Println() function no need to define data type
	fmt.Println(a)

	//Printing The datatype of the variable
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
