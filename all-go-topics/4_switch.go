package main

import "runtime"

func main(){
	val:=111

	switch{
	case val >=100:
		println("good")
	case val >=80:
		println("okay")
	default:
		println("not good")
	}

	switch os := runtime.GOOS; os{
	case "linux":
		println("linux")
	default:
		println("except linux")
	}

}