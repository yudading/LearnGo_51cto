package main

import "fmt"

func main() {
	var a int = 0

	switch a {

	case 0:
		fmt.Println("a is equals 0")
		fmt.Println("yes")
	case 10:
		fmt.Println("a is equals 10")
	default:
		fmt.Println("a is equals default")
	}

}
