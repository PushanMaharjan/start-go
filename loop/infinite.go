package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("infinite", i)
		i += 1;
	}
}
