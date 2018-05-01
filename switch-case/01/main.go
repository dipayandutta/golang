package main

import "fmt"

func main() {
	switch "Hello" {
	case "James":
		fmt.Println("Hello James")
	case "Daniel":
		fmt.Println("Hello Daniel")
	default:
		fmt.Println("Hello All!")
	}
}
