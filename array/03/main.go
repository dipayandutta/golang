package main

import "fmt"

func main() {
	var str [10]string
	fmt.Println(len(str))
	str[1] = "a"
	str[2] = "Hello World!"
	fmt.Println(str[1] + str[2])
}
