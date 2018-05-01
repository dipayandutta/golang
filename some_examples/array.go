package main 

import "fmt"

func main(){

	var a[5] int
	fmt.Println("emp:",a)
	fmt.Println("Length :",len(a))

	a = [5]int{0,1,2,3,4}

	fmt.Println("Third Element : ",a[3])
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	var twoD [2][3] int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D:",twoD)
}