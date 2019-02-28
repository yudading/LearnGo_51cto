package main

import "fmt"

func main() {
	const LENGHT int = 10
	const WIDHT int = 5

	var area int

	area = LENGHT * WIDHT

	fmt.Printf("面积为:%d", area)
	fmt.Println()

	const a, b, c = 1, false, "str"
	fmt.Println(a, b, c)
}
