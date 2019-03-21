package main

import "fmt"

func modify(p *int) {
	fmt.Println(p)
	*p = 100
}
func main() {
	var a int = 10
	modify(&a)
	fmt.Println(a)

	var p *int
	p = &a
	fmt.Println(*p)

	*p = 1000
	fmt.Println(a)

	var b int = 999
	p = &b
	*p = 5

	fmt.Println(a)
	fmt.Println(b)
}
