package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)
}
