package main

import "fmt"

func main() {

	var a [10]int
	var b [5]int

	fmt.Printf("len(a)=%d,len(b)=%d\n", len(a), len(b))

	a[0] = 1
	i := 1
	a[i] = 2

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}
}
