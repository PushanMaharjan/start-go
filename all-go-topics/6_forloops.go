package main

import "fmt"

func main(){
	ages := []int{11,27,8,10,17}

	for i:=0; i<len(ages); i++{
		fmt.Println(ages[i])
	}

	for i:=0; i<len(ages);{
		fmt.Println(ages[i])
		i++
	}

	for i:=0; i<10; i++{
		if i%2 !=0 {
			continue
		}
		fmt.Println(i)
	}

	for index,value := range ages{
		fmt.Println(index)
		fmt.Println(value)
	}
}