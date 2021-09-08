package main

import "fmt"

type Employee interface{
	GetName() string
}

type User struct{
	Name string
}

func (e *User) GetName() string {
	return e.Name
}

func PrintDetails (e Employee) {
	fmt.Println(e.GetName())
}

func main(){
	user:= &User{Name:"John wick"}
	PrintDetails(user)
}