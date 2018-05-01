package main

import "fmt"

func main() {
	n := average(11, 12.3, 30, 23, 1, 3, 44, 5)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	for i := 0; i < len(sf); i++ {
		fmt.Println(sf[i])

	}
	return total / float64(len(sf))
}
