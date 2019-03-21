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
}
