package main 

import (
	"fmt"
	"time"
)

func main(){

	i :=2
	fmt.Println("Write ",i , "as")

	switch i{

	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println("Today --> ",time.Now().Weekday())

	switch time.Now().Weekday(){

	case time.Saturday , time.Sunday:
		fmt.Println("It is a weekend")


	default :
		fmt.Println("It is working day")
	}

	t := time.Now()
	fmt.Println("Current System Time is ",t)

	
	whatAmI:=func(i interface{}){

		switch t := i.(type){
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hello")
}