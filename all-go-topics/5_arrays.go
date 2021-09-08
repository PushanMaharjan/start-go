package main

import "fmt"

func main(){
	arr := [8]int{1,2,3,4,5,6,7,8}

	fmt.Println(arr)

	// verbose
	var arr1 [8]string
	arr1[0]="1"

	fmt.Println(arr1)

	// slice
	arrSlice :=[]int{1,2,3,4,5,6,7,8}
	fmt.Println(arrSlice)

	//verbose
	var arrVerboseSlice []string
	arrVerboseSlice=append(arrVerboseSlice,"whatsup")
	arrVerboseSlice=append(arrVerboseSlice,"all good")
	fmt.Println(arrVerboseSlice)

}