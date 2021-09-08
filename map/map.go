package main

import "fmt"

func main() {
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	//    }

	colors["red"]="#ff0000"
	colors["green"]="#00ff00"
	colors["blue"]="#0000ff"

	fmt.Println(colors["red"])
	fmt.Println(colors["green"])
	fmt.Println(colors["blue"])
}