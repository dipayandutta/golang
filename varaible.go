package main 

import "fmt"
import "reflect"

func main(){

	var a int ;
	a = 10;

	fmt.Println("value of variable a is ",a)

	b := 10
	c := "Hello World"
	d := 12.3

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(reflect.TypeOf(b))

}