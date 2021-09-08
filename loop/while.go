package main

import "fmt"

func main() {
	sum := 1
	// while = for
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
