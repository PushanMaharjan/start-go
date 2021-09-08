package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:5]
	fmt.Println(s)

	fmt.Println(s[1:4])
	fmt.Println(s[:3])
	fmt.Println(s[2:])
	fmt.Println(s[:])

}