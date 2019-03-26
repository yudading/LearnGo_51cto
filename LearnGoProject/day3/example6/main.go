package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5}

	modify(a)
	fmt.Println("main a=", a)
}

func modify(a [6]int) {
	a[0] = 66666
	fmt.Println("modify a=", a)
}
