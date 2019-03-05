package main

import "fmt"

func main() {
	var a int = 50
	var b bool
	c := 'a'

	fmt.Printf("%v\n", a)
	fmt.Printf("%#v", b)
	fmt.Println()
	fmt.Printf("%T", c)
	fmt.Println()
	fmt.Printf("90%%")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100)
}
