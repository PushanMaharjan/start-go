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

// gets struct

func (e User) Print() {
	fmt.Println("User info")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Project: %s\n", e.Project.Name)
}

// manipulate (pointer)
func (e *User) AddAge() {
	e.Age++
}

// return
func (e *User) GetAge() int{
	return e.Age
} 


func main(){
	user := User{
		Name:"john wick", 
		Age: 22 , 
		Project: ProjectStruct{
			Name:"Baba Yaga", 
			time:11}}

	fmt.Printf("current: %d\n",user.GetAge())
	
	user.AddAge()		
	user.Print()

}