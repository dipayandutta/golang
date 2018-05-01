package main 

import "fmt"

func main(){

	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30

	fmt.Println(m)
	fmt.Println(m["key1"]+m["key2"])
	fmt.Println(len(m))


	n := map[string]int {"foo":12,"bar":13}
	fmt.Println(n)
	fmt.Println(n["foo"])

	for key := range n {
		fmt.Println("key[%s]",key)
	}

	for value := range n {

		fmt.Println("value[%s] ",value)
	}

	for key,value := range n{
		fmt.Println("key[%s] , value[%s]",key,value)
	}
}