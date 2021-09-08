package main

import "fmt"

type testStruct struct {
	first int
	second int
}

func main() {
	//struct actions

	val := testStruct{11, 27}

	fmt.Println(val)
	val.first=111
	fmt.Println(val)
	fmt.Println(val.first)

	val1 := testStruct{1111, 1111}  
	val2 := testStruct{first: 11111}  
	val3 := testStruct{}      
	valType  := &testStruct{11, 2}

	fmt.Println(val1, val2, val3, valType)

	//pointer in struct

	p := &val

	p.second = 270
	fmt.Println(val)
}
