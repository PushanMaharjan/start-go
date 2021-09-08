package main

import (
	"fmt"
)



func Sqrt(x float64) float64 {
	z := 1.0
	check := 0.0
	for check - z != 0 { 
		z = z - ( ((z*z) - x) / (2*z) )
		check = z - ( ((z*z) - x) / (2*z) )
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}