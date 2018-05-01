package main

import "fmt"

// Prints only oddd numbers using continue and break statement inside for loop
func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
