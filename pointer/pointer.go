package main 

import "fmt"

func main(){
	i:=11

	// assign pointer
	p:=&i 

	// read pointer
	fmt.Println(*p)

	//manipulate
	*p=11+11

	//output
	fmt.Println(*p)
	fmt.Println(i)
}