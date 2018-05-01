package main

import "fmt"

func zero(z int) {
	fmt.Println("Inside the zero function ", z)
	z = 1
	fmt.Println("Now z is ", z)

}
func main() {
	x := 5
	zero(x)
	fmt.Println("Inside the main function ", x)

}
