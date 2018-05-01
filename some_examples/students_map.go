package main

import "fmt"

func main(){

	student := map[string]map[string]int{
		"dipayan":map[string]int {"cs":50,"maths":60,"bengali":40},
		"kaushik":map[string]int {"cs":70,"maths":100,"bengali":56},
	}
/*
	//fmt.Println(student)
	//fmt.Println(student["dipayan"])
	//fmt.Println(student["kaushik"])

	var dipayanNumbers []int 
	for i := range student{
		//fmt.Println(i,"==>",student[i])
		for j := range student[i]{
			fmt.Println("\n",i,j,student[i][j])
			//fmt.Println("\n",student[i][j])
			if (i == "dipayan"){
				fmt.Println(student[i][j])
				dipayanNumbers := append(student[i][j])
			}
		}
	}
	fmt.Println(dipayanNumbers)

*/
	var totalDipayan,totalKaushik int

	for key := range student{
		fmt.Println(key)
		for _,value := range student[key]{
		fmt.Println(value)
		if (key=="dipayan"){
			totalDipayan += value
		}
		if(key=="kaushik"){
			totalKaushik += value
		}
	}
	}
	
	fmt.Println("Dipayan's Total Number",totalDipayan)
	fmt.Println("kaushik's Total Number",totalKaushik)
}