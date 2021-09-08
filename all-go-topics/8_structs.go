package main

import "fmt"


type ProjectStruct struct {
	Name string
	time int
}

type User struct{
	Name string
	Age int
	Project ProjectStruct
}


func main(){
	user := User{
		Name:"john wick", 
		Age: 22 , 
		Project: ProjectStruct{
			Name:"Baba Yaga", 
			time:11}}

	fmt.Printf("%+v\n", user)
	fmt.Println(user.Name)
	fmt.Println(user.Project.Name)
}