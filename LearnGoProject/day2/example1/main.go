package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	c := add
	fmt.Println(c)
}
