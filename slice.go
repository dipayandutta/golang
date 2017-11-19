package main 

import "fmt"

func main(){

	s := make([]string , 3)
	fmt.Println(s)
	fmt.Println(len(s))

	s[0] = "Hello"
	s[1] = "Beautiful" 
	s[2] = "World"

	fmt.Println(s)
	fmt.Println(s[1:3])

	s = append(s,"d")
	fmt.Println(s)
	fmt.Println(len(s))

	c := make([]string , len(s))
	copy(c,s)

	fmt.Println(c)

	sl := s[2:5]

	fmt.Println(sl)

	t := []string{"g","h","i"}
	fmt.Println(t)

	fmt.Println(t[2])
	fmt.Println(t[1:])

	twoD := make([][] int ,3)

	for i :=0;i<3 ;i++{
		innerLen := i+1
		twoD[i] = make([]int , innerLen)
			for j:=0;j< innerLen ; j++{
				twoD[i][j] = i+j
			}
		

	}
	fmt.Println("2d: ",twoD)
}