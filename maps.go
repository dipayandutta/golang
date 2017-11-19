package main 

import "fmt"

func main(){

	fruits := map[string]map[string]string {
		"Apple":map[string]string{"India":"10kg","USA":"20KG"},
		"Orange":map[string]string{"India":"50KG","China":"5KG"},
	}

	fmt.Println(fruits)
	fmt.Println(fruits["Apple"])
	fmt.Println(fruits["Apple"]["India"])

	for i := range fruits{
		fmt.Println(i,fruits[i])
		for j := range(fruits[i]) {
		fmt.Printf(" (%s) %s => %s\n", i, j, fruits[i][j])
	}
	}
}