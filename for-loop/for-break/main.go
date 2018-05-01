package main

import "fmt"

func zero(i int) {

	fmt.Println("Inside the function i is ==> ", i)
	i = 1
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}
}
func main() {

	var i int = 0
	zero(i)
}
