package main

import "fmt"

const metersToYard float64 = 1.09361

func main() {
	var meters float64
	fmt.Printf("Enter Meters")
	fmt.Scan(&meters)
	yards := meters * metersToYard
	fmt.Println(yards)
}
