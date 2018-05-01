package main

import "fmt"

var a = "This is a String"
var b, c string = "The B String", "The C String"
var d string

func main() {
	d = "The D String"
	var a, b int = 10, 20
	var c bool = true
	d := 12.3
	e, f := false, 12.4
	g := "G"
	h := `H`

	fmt.Println("A =", a)
	fmt.Println("B = ", b)
	fmt.Println("C =", c)
	fmt.Println("D = ", d)
	fmt.Println("E = ", e)
	fmt.Println("F = ", f)
	fmt.Println("G = ", g)
	fmt.Println("H =", h)

}
